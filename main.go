package main

import (
	"bwizz/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
  configs.LoadEnv()
  configs.ConnectPostgresSQL()
  
  app := fiber.New()
  app.Use(cors.New())

  api := app.Group("/api")

  api.Route("/quiz", func(router fiber.Router) {
    router.Get("/:id", func(c *fiber.Ctx) error {
      return c.JSON(":id")
    })
  })

  app.Get("/", index)

  app.Listen(":3000")
}

func index(c *fiber.Ctx) error {
  return c.SendString("Hello world")
}