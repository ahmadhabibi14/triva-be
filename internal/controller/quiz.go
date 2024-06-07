package controller

import (
	"triva/helper"
	"triva/internal/service"

	"github.com/gofiber/fiber/v2"
)

type QuizController struct {
	QuizPrefix  string
	quizService *service.QuizService
}

func NewQuizController(qs *service.QuizService) *QuizController {
	return &QuizController{
		QuizPrefix:  `/quiz`,
		quizService: qs,
	}
}

const (
	GetQuizzesAction = `/quizzes`
)

func (qc *QuizController) GetQuizzes(c *fiber.Ctx) error {
	quizzes, err := qc.quizService.GetQuizzes()
	if err != nil {
		response := helper.NewHTTPResponse(err.Error(), ``)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := helper.NewHTTPResponse(``, quizzes)
	return c.Status(fiber.StatusOK).JSON(response)
}
