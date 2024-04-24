package controller

import (
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (qc *AuthController) Login(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{`ok`: true})
}