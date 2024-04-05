package main

import (
	"bwizz/configs"
	"log"

	"github.com/gofiber/contrib/websocket"
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
  app.Use("/ws", func(c *fiber.Ctx) error {
    // IsWebSocketUpgrade returns true if the client
    // requested upgrade to the WebSocket protocol.
    if websocket.IsWebSocketUpgrade(c) {
        c.Locals("allowed", true)
        return c.Next()
    }
    return fiber.ErrUpgradeRequired
})

app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
    // c.Locals is added to the *websocket.Conn
    log.Println(c.Locals("allowed"))  // true
    log.Println(c.Params("id"))       // 123
    log.Println(c.Query("v"))         // 1.0
    log.Println(c.Cookies("session")) // ""

    // websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
    var (
        mt  int
        msg []byte
        err error
    )
    for {
        if mt, msg, err = c.ReadMessage(); err != nil {
            log.Println("read:", err)
            break
        }
        log.Printf("recv: %s", msg)

        if err = c.WriteMessage(mt, msg); err != nil {
            log.Println("write:", err)
            break
        }
    }

}))

  app.Listen(":3000")
}

func index(c *fiber.Ctx) error {
  return c.SendString("Hello world")
}