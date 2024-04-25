package controller

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	AUTH_COOKIE = `session_id`
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (qc *AuthController) Login(ctx *fiber.Ctx) error {
	qc.setCookie(ctx, `---`)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{`ok`: true})
}

func (qc *AuthController) Register(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{`ok`: true})
}

func (qc *AuthController) ResetPassword(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{`ok`: true})
}

func (qc *AuthController) ForgotPassword(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{`ok`: true})
}

func (qc *AuthController) setCookie(ctx *fiber.Ctx, sessionId string) {
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
