package internal

import (
	"log"
	"os"
	"time"
	"triva/configs"
	"triva/internal/bootstrap/database"
	"triva/internal/bootstrap/logger"
	"triva/internal/controller"
	"triva/internal/service"
	"triva/internal/web"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	// Common
	httpServer *fiber.App
	db 				 *database.Database
	// Services
	authService *service.AuthService
	userService *service.UserService
	quizService *service.QuizService
	netService  *service.NetService
}

func NewApp() *App {
	return &App{}
}

func (a *App) Init() {
	a.setupEnv()
	a.setupLogger()
	a.setupDatabases()
	a.setupServices()
	a.setupHTTP()
	a.apiDocs()

	port := ":" + os.Getenv("WEB_PORT")
	if port == ":" { port = ":3000" }
	log.Fatal(a.httpServer.Listen(port))
}

func (a *App) setupHTTP() {
	app := web.NewWebserver()
	middleware := web.NewMiddlewares(app, a.db)
	middleware.Init()

	authController := controller.NewAuthController(a.authService)
	app.Route(authController.AuthPrefix, func(router fiber.Router) {
		router.Post(controller.LoginAction, authController.Login)
		router.Post(controller.RegisterAction, authController.Register)
	})

	userController := controller.NewUserController(a.userService)
	app.Route(userController.UserPrefix, func(router fiber.Router) {
		router.Post(controller.UpdateAvatarAction, middleware.OPT_Auth, userController.UpdateAvatar)
	})

	quizController := controller.NewQuizController(a.quizService)
	app.Route(quizController.QuizPrefix, func(router fiber.Router) {
		router.Get(controller.GetQuizzesAction, middleware.OPT_Auth, quizController.GetQuizzes)
	})

	gameController := controller.NewGameController(a.netService)
	app.Route(gameController.GamePrefix, func(router fiber.Router) {
		router.Get(controller.HostAction, middleware.OPT_Auth, middleware.OPT_WebSocket, websocket.New(gameController.Host))
		router.Get(controller.PlayerAction, middleware.OPT_WebSocket, websocket.New(gameController.Player))
	})

	a.httpServer = app
}

func (a *App) setupServices() {
	a.authService = service.NewAuthService(a.db)
	a.userService = service.NewUserService(a.db)
	a.quizService = service.NewQuizService(a.db)
	a.netService = service.NewNetService(a.quizService, a.db)
}

func (a *App) setupDatabases() {
	pq := configs.ConnectPostgresSQL()

	rd := configs.NewRedisClient()
	_, err := rd.Ping().Result()
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

func (a *App) apiDocs() {
	a.httpServer.Use(swagger.New(swagger.Config{
		BasePath: `/`,
		FilePath: `./docs/swagger.json`,
		Path: `docs`,
		Title: `Triva API Documentation`,
		CacheAge: int(time.Minute) * 30,
	}))
}