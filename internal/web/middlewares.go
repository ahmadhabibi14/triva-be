package web

import (
	"bwizz/helper"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Middlewares struct {
	app *fiber.App
}

func NewMiddlewares(app *fiber.App) *Middlewares { return &Middlewares{app: app} }

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
	m.app.Use(logger.New(logger.Config{
		Format:     "${time} | ${status} | ${latency} | ${method} | ${path}\n",
		TimeFormat: "2006/01/02 03:04 PM",
		TimeZone:   "Asia/Makassar",
		Output:     os.Stdout,
	}))
}