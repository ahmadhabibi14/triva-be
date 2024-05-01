package web

import (
	"os"
	"time"
	"triva/configs"
	"triva/helper"
	"triva/internal/controller"

	"github.com/go-redis/redis"
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
	rd  *redis.Client
}

func NewMiddlewares(app *fiber.App, log *zerolog.Logger, rd *redis.Client) *Middlewares {
	return &Middlewares{
		app: app,
		log: log,
		rd:  rd,
	}
}

func (m *Middlewares) Init() {
	m.RateLimiter()
	m.Cors()
	m.Logger()
	m.Recover()
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

func (m *Middlewares) Auth() {
	m.app.Use(func(c *fiber.Ctx) error {
		sessionId := c.Cookies(controller.AUTH_COOKIE, ``)
		apiKey := c.Get("X-API-KEY", ``)

		// TODO
		if len(sessionId) == 0 && len(apiKey) == 0 {
			return c.Next()
		}
		return c.Next()
	})
}
