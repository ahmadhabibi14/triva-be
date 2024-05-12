package controller

import (
	"os"
	"time"
	"triva/configs"
	"triva/helper"
	"triva/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthPrefix  string
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		AuthPrefix:  `/api/auth`,
		authService: authService,
	}
}

type (
	LoginIn struct {
		Username string `json:"username" form:"username" validate:"required,omitempty,min=5"`
		Password string `json:"password" form:"password" validate:"required,min=8"`
	}
	LoginOut struct {
		Session string `json:"session" form:"session"`
		Message string `json:"message" form:"message"`
	}
)

const (
	LoginAction = `/login`
	LoginOkMsg  = `login successful !`
)

func (ac *AuthController) Login(c *fiber.Ctx) error {
	in, err := helper.ReadJSON[LoginIn](c, c.Body())
	if err != nil {
		response := helper.NewHTTPResponse(fiber.StatusBadRequest, err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	sessionKey, err := ac.authService.Login(in.Username, in.Password)
	if err != nil {
		response := helper.NewHTTPResponse(fiber.StatusBadRequest, err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	ac.setCookie(c, sessionKey)

	out := LoginOut{Session: sessionKey, Message: LoginOkMsg}
	response := helper.NewHTTPResponse(fiber.StatusCreated, ``, out)
	return c.Status(fiber.StatusOK).JSON(response)
}

type (
	RegisterIn struct {
		Username string `json:"username" form:"username" validate:"required,omitempty,min=5"`
		FullName string `json:"full_name" form:"full_name" validate:"required,omitempty,min=5"`
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required,min=8"`
	}
	RegisterOut struct {
		Message string `json:"message" form:"message"`
	}
)

const (
	RegisterAction = `/register`
	RegisterOkMsg  = `user created !`
)

func (ac *AuthController) Register(c *fiber.Ctx) error {
	in, err := helper.ReadJSON[RegisterIn](c, c.Body())
	if err != nil {
		response := helper.NewHTTPResponse(fiber.StatusBadRequest, err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	err = ac.authService.Register(
		in.Username, in.FullName, in.Email, in.Password,
	)

	if err != nil {
		response := helper.NewHTTPResponse(fiber.StatusBadRequest, err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	out := RegisterOut{Message: RegisterOkMsg}
	response := helper.NewHTTPResponse(fiber.StatusCreated, ``, out)
	return c.Status(fiber.StatusCreated).JSON(response)
}

func (ac *AuthController) ResetPassword(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{`ok`: true})
}

func (ac *AuthController) ForgotPassword(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{`ok`: true})
}

func (ac *AuthController) setCookie(c *fiber.Ctx, sessionId string) {
	// 2 months expired
	expiration := time.Now().AddDate(0, 2, 0)

	c.Cookie(&fiber.Cookie{
		Name:     configs.AUTH_COOKIE,
		Value:    sessionId,
		Expires:  expiration,
		SameSite: `Lax`,
		Secure:   os.Getenv(`WEB_ENV`) == `prod`,
		HTTPOnly: true,
	})
}
