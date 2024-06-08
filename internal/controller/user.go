package controller

import (
	"time"
	"triva/helper"
	"triva/internal/repository/users"
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

type (
	UpdateAvatarIn struct {
		Avatar 			string 			`json:"avatar" form:"avatar"`
	}
	UpdateAvatarOut struct {
		Id 					string 		`json:"id"`
		Username 		string 		`json:"username"`
		FullName 		string 		`json:"full_name"`
		Email 			string 		`json:"email"`
		AvatarURL		string 		`json:"avatar_url"`
		GoogleId 		string 		`json:"google_id"`
		FacebookId	string 		`json:"facebook_id"`
		GithubId 		string 		`json:"github_id"`
		CreatedAt 	time.Time `json:"created_at"`
		UpdatedAt 	time.Time `json:"updated_at"`
	}
)

const (
	UpdateAvatarAction = `/update-avatar`
	UpdateAvatarOkMsg  = `avatar updated`
)

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

	out := UpdateAvatarOut{
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