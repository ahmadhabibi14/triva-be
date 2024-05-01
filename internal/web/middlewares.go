package web

import (
	"os"
	"time"
	"triva/configs"
	"triva/helper"
	"triva/internal/repository/users"

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
		KeyGenerator: func(ctx *fiber.Ctx) string {
			return ctx.IP()
		},
		LimitReached: func(ctx *fiber.Ctx) error {
			var errMessage string = "You have exceeded your rate limit. Please try again a few minutes later."

			response := helper.NewHTTPResponse(fiber.StatusTooManyRequests, errMessage, "")
			return ctx.Status(fiber.StatusTooManyRequests).JSON(response)
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
		StackTraceHandler: func(ctx *fiber.Ctx, e interface{}) {
			m.log.Error().Str("path", ctx.Path()).Err(e.(error)).Msg("received unexpected panic error")
		},
	}))
}


// Optional Middlewares

const (
	errMsgUnauthorized	= `you are unauthorized to process this operation`
	errMsgInvalidKey		= `invalid session key`
)

func (m *Middlewares) OPT_Auth(ctx *fiber.Ctx) error {
	sessionId := ctx.Cookies(configs.AUTH_COOKIE, ``)
	apiKey := ctx.Get("X-API-KEY", ``)

	var KEY string = sessionId

	if sessionId == `` {
		KEY = apiKey
		if apiKey == `` {
			response := helper.NewHTTPResponse(fiber.StatusUnauthorized, errMsgUnauthorized)
			return ctx.Status(fiber.StatusUnauthorized).JSON(response)
		}
	}

	session := users.NewSessionMutator(m.rd)
	err := session.GetSession(users.SESSION_PREFIX + KEY)
	
	if err != nil {
		m.log.Error().Str("error", err.Error()).Msg("cannot get session data for " + KEY)

		ctx.ClearCookie(configs.AUTH_COOKIE)
		response := helper.NewHTTPResponse(fiber.StatusUnauthorized, errMsgInvalidKey)
		return ctx.Status(fiber.StatusUnauthorized).JSON(response)
	}

	return ctx.Next()
}
