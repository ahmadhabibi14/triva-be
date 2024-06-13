package controller

import (
	"triva/helper"
	"triva/internal/repository/users"
	"triva/internal/response"
	"triva/internal/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserPrefix  string
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		UserPrefix: `/user`,
		userService: userService,
	}
}

const (
	UpdateAvatarAction = `/update-avatar`
	UpdateAvatarOkMsg  = `avatar updated`
)

// @Summary 			Update avatar
// @Tags					User
// @Accept				mpfd
// @Param 				requestBody  body  UpdateAvatarIn  true  "Avatar image"
// @Success				200 {object} UpdateAvatarOut
// @Produce				json
// @Router				/user/updateAvatar [post]
func (uc *UserController) UpdateAvatar(c *fiber.Ctx, session *users.Session) error {
	imgFile, err := c.FormFile("avatar")
	if err != nil {
		response := helper.NewHTTPResponse(`file cannot be empty`, nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	user, err := uc.userService.UpdateAvatar(imgFile, session.UserID)
	if err != nil {
		response := helper.NewHTTPResponse(err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	out := response.UpdateAvatarOut{
		Id: user.Id,
		Username: user.Username,
		FullName: user.FullName,
		Email: user.Email,
		AvatarURL: user.AvatarURL,
		GoogleId: user.GoogleId,
		FacebookId: user.FacebookId,
		GithubId: user.GithubId,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	response := helper.NewHTTPResponse(``, out)
	return c.Status(fiber.StatusOK).JSON(response)
}