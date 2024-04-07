package internal

import (
	"bwizz/configs"
	"bwizz/helper"
	"bwizz/internal/controller"
	"bwizz/internal/repository/quizzes"
	"bwizz/internal/service"
	"errors"
	"log"
	"os"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
)

type App struct{
	httpServer *fiber.App
	db *sqlx.DB

	quizService *service.QuizService
	netService *service.NetService
}

func (a *App) Init() {
	a.setupEnv()
	a.setupDB()
	a.setupServices()
	a.setupHTTP()

	log.Fatal(a.httpServer.Listen(":3000"))
}

func (a *App) setupHTTP() {
	app := fiber.New(fiber.Config{
		AppName: os.Getenv("PROJECT_NAME"),
		Prefork: false,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			var code int = fiber.StatusNotFound
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			resp := helper.NewHTTPResponse(code, e.Error(), nil)

			c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
			return c.Status(code).JSON(resp)
		},
	})
	app.Use(cors.New())

	quizController := controller.NewQuizController(a.quizService)
	app.Get("/api/quizzes", quizController.GetQuizzes)

	wsController := controller.NewWebsocketController(a.netService)
	app.Get("/ws", websocket.New(wsController.WS))

	a.httpServer = app
}

func (a *App) setupServices() {
	a.quizService = service.NewQuizService(quizzes.NewQuizMutator())
	a.netService = service.NewNetService(a.quizService)
}

func (a *App) setupDB() {
	err := configs.ConnectPostgresSQL()
	if err != nil {
    panic("failed to connect to database")
  }
	a.db = configs.PostgresDB
}

func (a *App) setupEnv() {
	envFilePath := ".env"
  err := configs.LoadEnv(envFilePath)
	if err != nil {
		panic("cannot load "+envFilePath)
	}
}