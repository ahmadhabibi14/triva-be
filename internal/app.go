package internal

import (
	"log"
	"os"
	"triva/configs"
	"triva/internal/bootstrap/database"
	"triva/internal/bootstrap/logger"
	"triva/internal/controller"
	"triva/internal/service"
	"triva/internal/web"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	httpServer *fiber.App
	db 				 *database.Database

	authService *service.AuthService
	quizService *service.QuizService
	netService  *service.NetService
}

func (a *App) Init() {
	a.setupEnv()
	a.setupLogger()
	a.setupDatabases()
	a.setupServices()
	a.setupHTTP()

	log.Fatal(a.httpServer.Listen(":" + os.Getenv("WEB_PORT")))
}

func (a *App) setupHTTP() {
	app := web.NewWebserver()
	middleware := web.NewMiddlewares(app, a.db)
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
	app.Get(gameController.GameAction, websocket.New(gameController.Game))

	a.httpServer = app
}

func (a *App) setupServices() {
	a.authService = service.NewAuthService(a.db)
	a.quizService = service.NewQuizService(a.db)
	a.netService = service.NewNetService(a.quizService, a.db)
}

func (a *App) setupDatabases() {
	pq, err := configs.ConnectPostgresSQL()
	if err != nil {
		logger.Log.Panic().Str("error", err.Error()).Msg("failed to connect to database")
	}

	rd := configs.NewRedisClient()
	_, err = rd.Ping().Result()
	if err != nil {
		logger.Log.Panic().Str("error", err.Error()).Msg("failed to connect redis")
	}

	db := database.NewDatabase(pq, rd)

	a.db = db
}

func (a *App) setupEnv() {
	configs.LoadEnv()
}

func (a *App) setupLogger() {
	logger.InitLogger()
}
