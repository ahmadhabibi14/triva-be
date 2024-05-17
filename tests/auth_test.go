package tests

import (
	"bytes"
	"io"
	"net/http/httptest"
	"testing"
	"time"
	"triva/internal/controller"
	"triva/internal/service"
	"triva/internal/web"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	authService := service.NewAuthService(DB)
	authController := controller.NewAuthController(authService)

	app := web.NewWebserver()
	middleware := web.NewMiddlewares(app, LOG, DB)
	middleware.Init()

	var (
		loginPath    string = authController.AuthPrefix + controller.LoginAction
		registerPath string = authController.AuthPrefix + controller.RegisterAction
	)

	// Routes
	app.Post(loginPath, authController.Login)
	app.Post(registerPath, authController.Register)

	// MISC config
	timeOut := 10 * time.Second

	t.Run(`authLogin`, func(t *testing.T) {
		payload := controller.LoginIn{
			Username: "habibi",
			Password: "habibi12345678",
		}

		jsonPayload, err := json.Marshal(payload)
		assert.Nil(t, err, `failed to marshal payload`)

		req := httptest.NewRequest(fiber.MethodPost, loginPath, bytes.NewBuffer(jsonPayload))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

		resp, err := app.Test(req, int(timeOut))
		assert.Nil(t, err, `failed to make request`)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode, `failed to make request`)

		body, err := io.ReadAll(resp.Body)
		assert.Nil(t, err, `failed to read body`)

		t.Log(`response body:`, string(body))
	})
}
