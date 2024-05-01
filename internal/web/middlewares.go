package web

import (
	"os"
	"time"
	"triva/configs"
	"triva/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"
)

type Middlewares struct {
	app *fiber.App
	log *zerolog.Logger
}

func NewMiddlewares(app *fiber.App, log *zerolog.Logger) *Middlewares {
	return &Middlewares{
		app: app}
}

func (m *Middlewares) Init() {
	m.RateLimiter()
	m.Cors()
	m.Logger()
}

func (m *Middlewares) RateLimiter() {
	m.app.Use(limiter.New(limiter.Config{
		Max:        300,
		Expiration: 2 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			var errMessage string = "You have exceeded your rate limit. Please try again a few minutes later."

			response := helper.NewHTTPResponse(fiber.StatusTooManyRequests, errMessage, "")
			return c.Status(fiber.StatusTooManyRequests).JSON(response)
		},
	}))
}

func (m *Middlewares) Cors() {
	m.app.Use(cors.New())
}

func (m *Middlewares) Logger() {
	var conf logger.Config

	if os.Getenv("WEB_ENV") == `prod` {
		file, _ := os.OpenFile(
			configs.PATH_WEBACCESS_LOG, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666,
		)
		conf = logger.Config{
			Format:        "{\"time\": \"${time}\", \"status\": \"${status}\", \"ip\": \"${ip}\", \"ips\": \"${ips}\", \"latency\": \"${latency}\", \"method\": \"${method}\", \"path\": \"${path}\"\n",
			TimeFormat:    "2006-01-02T03:00:55+08:00",
			TimeZone:      "Asia/Makassar",
			Output:        file,
			DisableColors: true,
		}
	} else {
		conf = logger.Config{
			Format:     "${time} | ${status} | ${latency} | ${method} | ${path}\n",
			TimeFormat: "2006/01/02 03:04 PM",
			TimeZone:   "Asia/Makassar",
			Output:     os.Stdout,
		}
	}

	m.app.Use(logger.New(conf))
}

func (m *Middlewares) Recover() {
	m.app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			m.log.Error().Str("path", c.Path()).Err(e.(error)).Msg("received unexpected panic error")
		},
	}))
}
