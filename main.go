package main

import (
	"github.com/ahmadhabibi14/kahoot-clone/configs"
	"github.com/gofiber/fiber/v2"
)

func main() {
  configs.LoadEnv()
  configs.ConnectPostgresSQL()
  
  app := fiber.New()

  app.Get("/", index)

  app.Listen(":3000")
}

func index(c *fiber.Ctx) error {
  return c.SendString("Hello world")
}