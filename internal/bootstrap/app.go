package bootstrap

import (
	"os"
	"os/signal"
	"syscall"
	"time"
	"triva/configs"
	"triva/helper"
	"triva/internal/bootstrap/database"
	"triva/internal/bootstrap/logger"
	"triva/internal/bootstrap/web"
	"triva/internal/controller"
	"triva/internal/service"

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

func (a *App) Run() {
	waits := make(chan int)

	a.setupEnv()
	a.setupLogger()
	a.setupDatabases()
	a.setupServices()
	a.setupHTTP()
	a.apiDocs()
	go a.shutdown()

	port := ":" + os.Getenv("WEB_PORT"); if port == ":" { port = ":3000" }

	if err := a.httpServer.Listen(port); err != nil {
		logger.Log.Err(err).Msg("failed to start http server")
		a.closeServices()
		os.Exit(1)
	}

	<-waits
}

func (a *App) setupHTTP() {
	app := web.NewWebserver()
	middleware := web.NewMiddlewares(app, a.db)
	middleware.Init()

	// serve static files
	app.Static("/", configs.OS_PATH_STATIC_FILES)

	authController := controller.NewAuthController(a.authService)
	app.Route(authController.AuthPrefix, func(router fiber.Router) {
		router.Post(controller.LoginAction, authController.Login)
		router.Post(controller.RegisterAction, authController.Register)
	})

	userController := controller.NewUserController(a.userService)
	app.Route(userController.UserPrefix, func(router fiber.Router) {
		router.Post(controller.UpdateAvatarAction, func(c *fiber.Ctx) error {
			session, err := web.GetSession(a.db, c)
			if err != nil {
				response := helper.NewHTTPResponse(err.Error(), nil)
				return c.Status(fiber.StatusBadRequest).JSON(response)
			}

			return userController.UpdateAvatar(c, session)
		})
	})

	quizController := controller.NewQuizController(a.quizService)
	app.Route(quizController.QuizPrefix, func(router fiber.Router) {
		router.Get(controller.GetQuizzesAction, func(c *fiber.Ctx) error {
			session, err := web.GetSession(a.db, c)
			if err != nil {
				response := helper.NewHTTPResponse(err.Error(), nil)
				return c.Status(fiber.StatusBadRequest).JSON(response)
			}

			return quizController.GetQuizzes(c, session)
		})
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
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path: "docs",
		Title: "Triva API Documentation",
		CacheAge: int(time.Minute) * 30,
	}))
}

func (a *App) shutdown() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		<-s

		logger.Log.Info().Msg("shutting down...")
		a.closeServices()

		os.Exit(0)
	}()
}

func (a *App) closeServices() {
	if err := a.httpServer.Shutdown(); err != nil {
		logger.Log.Err(err).Msg("failed to shutdown [httpserver]")
	} else {
		logger.Log.Info().Msg("cleaned up [httpserver]")
	}

	if err := a.db.DB.Close(); err != nil {
		logger.Log.Err(err).Msg("failed to shutdown [postgresql]")
	} else {
		logger.Log.Info().Msg("cleaned up [postgresql]")
	}

	if err := a.db.RD.Close(); err != nil {
		logger.Log.Err(err).Msg("failed to shutdown [redis]")
	} else {
		logger.Log.Info().Msg("cleaned up [redis]")
	}
}