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
		QuizPrefix:  `/api/quiz`,
		quizService: qs,
	}
}

const (
	GetQuizzesAction = `/quizzes`
)

func (qc *QuizController) GetQuizzes(ctx *fiber.Ctx) error {
	quizzes, err := qc.quizService.GetQuizzes()
	if err != nil {
		response := helper.NewHTTPResponse(fiber.StatusBadRequest, err.Error(), ``)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := helper.NewHTTPResponse(fiber.StatusOK, ``, quizzes)
	return ctx.Status(fiber.StatusOK).JSON(response)
}
