package internal

import (
	"log"
	"os"
	"triva/configs"
	"triva/internal/controller"
	"triva/internal/database"
	"triva/internal/service"
	"triva/internal/web"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type App struct {
	httpServer *fiber.App
	log        *zerolog.Logger
	db 				 *database.Database

	authService *service.AuthService
	quizService *service.QuizService
	netService  *service.NetService
}

func (a *App) Init() {
	a.setupEnv()
	a.setupLogger()
	a.setupDatabase()
	a.setupServices()
	a.setupHTTP()

	log.Fatal(a.httpServer.Listen(":" + os.Getenv("WEB_PORT")))
}

func (a *App) setupHTTP() {
	app := web.NewWebserver()
	middleware := web.NewMiddlewares(app, a.log, a.db)
	middleware.Init()

	authController := controller.NewAuthController(a.authService)
	app.Route(authController.AuthPrefix, func(app fiber.Router) {
		app.Post(controller.LoginAction, authController.Login)
		app.Post(controller.RegisterAction, authController.Register)
	})

	quizController := controller.NewQuizController(a.quizService)
	app.Route(quizController.QuizPrefix, func(app fiber.Router) {
		app.Get(controller.GetQuizzesAction, middleware.OPT_Auth, quizController.GetQuizzes)
	})

	gameController := controller.NewGameController(a.netService)
	app.Get("/game", websocket.New(gameController.Game))

	a.httpServer = app
}

func (a *App) setupServices() {
	a.authService = service.NewAuthService(a.db)
	a.quizService = service.NewQuizService(a.db)
	a.netService = service.NewNetService(a.quizService, a.db)
}

func (a *App) setupDatabase() {
	pq, err := configs.ConnectPostgresSQL()
	if err != nil {
		a.log.Panic().Str("error", err.Error()).Msg("failed to connect to database")
	}

	rd := configs.NewRedisClient()
	_, err = rd.Ping().Result()
	if err != nil {
		a.log.Panic().Str("error", err.Error()).Msg("failed to connect redis")
	}

	db := database.NewDatabase(pq, rd, a.log)

	a.db = db
}

func (a *App) setupEnv() {
	configs.LoadEnv()
}

func (a *App) setupLogger() {
	a.log = configs.NewLogger()
}
