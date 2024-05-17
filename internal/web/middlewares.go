package web

import (
	"os"
	"time"
	"triva/configs"
	"triva/helper"
	"triva/internal/database"
	"triva/internal/repository/users"

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
	db *database.Database
}

func NewMiddlewares(app *fiber.App, log *zerolog.Logger, db *database.Database) *Middlewares {
	return &Middlewares{
		app: app,
		log: log,
		db: db,
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
			var errMessage string = "you have exceeded your rate limit, please try again a few moments later"

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

// Optional Middlewares

const (
	errMsgUnauthorized	= `you are unauthorized to process this operation`
	errMsgInvalidKey		= `invalid session key`
)

func (m *Middlewares) OPT_Auth(c *fiber.Ctx) error {
	sessionId := c.Cookies(configs.AUTH_COOKIE, ``)
	apiKey := c.Get("X-API-KEY", ``)

	var KEY string = sessionId
	if sessionId == `` { KEY = apiKey }
	
	if KEY == `` {
		response := helper.NewHTTPResponse(fiber.StatusUnauthorized, errMsgUnauthorized, nil)
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	session := users.NewSessionMutator(m.db)
	
	if err := session.GetSession(KEY); err != nil {
		m.log.Error().Str("error", err.Error()).Msg("cannot get session data for " + KEY)

		c.ClearCookie(configs.AUTH_COOKIE)
		response := helper.NewHTTPResponse(fiber.StatusUnauthorized, errMsgInvalidKey, nil)
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	return c.Next()
}
