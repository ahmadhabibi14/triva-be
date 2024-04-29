package controller

import (
	"os"
	"time"
	"triva/helper"
	"triva/internal/service"

	"github.com/gofiber/fiber/v2"
)

const (
	AUTH_COOKIE = `session_id`
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ac *AuthController) Login(ctx *fiber.Ctx) error {
	ac.setCookie(ctx, `---`)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{`ok`: true})
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

func (ac *AuthController) Register(ctx *fiber.Ctx) error {
	in, err := helper.ReadJSON[RegisterIn](ctx, ctx.Body())
	if err != nil {
		response := helper.NewHTTPResponse(fiber.StatusBadRequest, err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	err = ac.authService.Register(
		in.Username, in.FullName, in.Email, in.Password,
	)

	if err != nil {
		response := helper.NewHTTPResponse(fiber.StatusBadRequest, err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	out := RegisterOut{Message: `user created`}
	response := helper.NewHTTPResponse(fiber.StatusCreated, ``, out)
	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (ac *AuthController) ResetPassword(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{`ok`: true})
}

func (ac *AuthController) ForgotPassword(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{`ok`: true})
}

func (ac *AuthController) setCookie(ctx *fiber.Ctx, sessionId string) {
	// 2 months expired
	expiration := time.Now().AddDate(0, 2, 0)

	ctx.Cookie(&fiber.Cookie{
		Name:     AUTH_COOKIE,
		Value:    sessionId,
		Expires:  expiration,
		SameSite: `Lax`,
		Secure:   os.Getenv(`WEB_ENV`) == `prod`,
		HTTPOnly: false,
	})
}
