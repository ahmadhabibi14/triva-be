package internal

import (
	"bwizz/configs"
	"bwizz/internal/controller"
	"bwizz/internal/repository/quizzes"
	"bwizz/internal/service"
	"bwizz/internal/web"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
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
	app := web.NewWebserver()
	middleware := web.NewMiddlewares(app)
	middleware.Init()

	quizController := controller.NewQuizController(a.quizService)
	app.Get("/api/quizzes", quizController.GetQuizzes)

	wsController := controller.NewWebsocketController(a.netService)
	app.Get("/ws", websocket.New(wsController.WS))

	a.httpServer = app
}

func (a *App) setupServices() {
	a.quizService = service.NewQuizService(quizzes.NewQuizMutator(a.db))
	a.netService = service.NewNetService(a.quizService)
}

func (a *App) setupDB() {
	db, err := configs.ConnectPostgresSQL()
	if err != nil {
    panic("failed to connect to database")
  }
	a.db = db
}

func (a *App) setupEnv() {
	envFilePath := ".env"
  err := configs.LoadEnv(envFilePath)
	if err != nil {
		panic("cannot load "+envFilePath)
	}
}