package controller

import (
	"bwizz/internal/service"

	"github.com/gofiber/fiber/v2"
)

type QuizController struct {
	quizService *service.QuizService
}

func NewQuizController(qs *service.QuizService) *QuizController {
	return &QuizController{quizService: qs}
}

func (qc *QuizController) GetQuizzes(ctx *fiber.Ctx) error {
	quizzes, err := qc.quizService.GetQuizzes()
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(quizzes)
}