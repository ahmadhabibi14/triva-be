package controller

import (
	"os"
	"time"
	"triva/configs"
	"triva/helper"
	"triva/internal/request"
	_ "triva/internal/response"
	"triva/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthPrefix  string
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		AuthPrefix:  `/auth`,
		authService: authService,
	}
}

const LoginAction = `/login`

// @Summary 			Login to authenticated
// @Tags					Auth
// @Param 				requestBody  body  request.LoginIn  true  "User credentials"
// @Success				200 {object} response.LoginOut "Login Out"
// @Produce				json
// @Router				/auth/login [post]
func (ac *AuthController) Login(c *fiber.Ctx) error {
	loginIn, err := helper.ReadBody[request.LoginIn](c, c.Body())
	if err != nil {
		response := helper.NewHTTPResponse(err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	loginOut, err := ac.authService.Login(loginIn)
	if err != nil {
		response := helper.NewHTTPResponse(err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	ac.setCookie(c, loginOut.SessionKey)

	response := helper.NewHTTPResponse(``, loginOut)
	return c.Status(fiber.StatusOK).JSON(response)
}

const RegisterAction = `/register`

// @Summary 			Regiser to create an account
// @Tags					Auth
// @Param 				requestBody  body  request.RegisterIn  true  "User data"
// @Success				200 {object} response.RegisterOut "Register Out"
// @Produce				json
// @Router				/auth/register [post]
func (ac *AuthController) Register(c *fiber.Ctx) error {
	registerIn, err := helper.ReadBody[request.RegisterIn](c, c.Body())
	if err != nil {
		response := helper.NewHTTPResponse(err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	registerOut, err := ac.authService.Register(registerIn)
	if err != nil {
		response := helper.NewHTTPResponse(err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := helper.NewHTTPResponse(``, registerOut)
	return c.Status(fiber.StatusCreated).JSON(response)
}

func (ac *AuthController) ResetPassword(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{`ok`: true})
}

func (ac *AuthController) ForgotPassword(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{`ok`: true})
}

// ##################################### //
// ########### Utilities ############### //
// ##################################### //

func (ac *AuthController) setCookie(c *fiber.Ctx, sessionKey string) {
	// 2 months expired
	expiration := time.Now().AddDate(0, 2, 0)

	c.Cookie(&fiber.Cookie{
		Name:     configs.AUTH_COOKIE,
		Value:    sessionKey,
		Expires:  expiration,
		SameSite: `Lax`,
		Secure:   os.Getenv(`PROJECT_ENV`) == `prod`,
		HTTPOnly: true,
	})
}
