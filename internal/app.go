package internal

import (
	"log"
	"os"
	"triva/configs"
	"triva/internal/controller"
	"triva/internal/repository/quizzes"
	"triva/internal/service"
	"triva/internal/web"

	"github.com/go-redis/redis"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type App struct{
	httpServer *fiber.App
	db *sqlx.DB
	rd *redis.Client
	log *zerolog.Logger

	authService *service.AuthService
	quizService *service.QuizService
	netService *service.NetService
}

func (a *App) Init() {
	a.setupEnv()
	a.setupLogger()
	a.setupDB()
	a.setupRedis()
	a.setupServices()
	a.setupHTTP()

	log.Fatal(a.httpServer.Listen(":"+os.Getenv("WEB_PORT")))
}

func (a *App) setupHTTP() {
	app := web.NewWebserver()
	middleware := web.NewMiddlewares(app)
	middleware.Init()

	app.Use(recover.New())

	authController := controller.NewAuthController(a.authService)
	app.Post("/api/auth/login", authController.Login)
	app.Post("/api/auth/register", authController.Register)

	quizController := controller.NewQuizController(a.quizService)
	app.Get("/api/quizzes", quizController.GetQuizzes)

	wsController := controller.NewWebsocketController(a.netService)
	app.Get("/ws", websocket.New(wsController.WS))

	a.httpServer = app
}

func (a *App) setupServices() {
	a.authService = service.NewAuthService(a.db, a.rd)
	a.quizService = service.NewQuizService(quizzes.NewQuizMutator(a.db))
	a.netService = service.NewNetService(a.quizService, a.db)
}

func (a *App) setupDB() {
	db, err := configs.ConnectPostgresSQL()
	if err != nil {
    a.log.Panic().Str("error", err.Error()).Msg("failed to connect to database")
  }
	a.db = db
}

func (a *App) setupRedis() {
	rd := configs.NewRedisClient()

	_, err := rd.Ping().Result()
	if err != nil {
		a.log.Panic().Str("error", err.Error()).Msg("failed to connect redis")
	}
	
	a.rd = rd
}

func (a *App) setupEnv() {
  configs.LoadEnv()
}

func (a *App) setupLogger() {
	a.log = configs.NewLogger()
}