package controller

import (
	"triva/helper"
	_ "triva/internal/repository/quizzes"
	"triva/internal/repository/users"
	"triva/internal/request"
	_ "triva/internal/response"
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

const GetQuizzesAction = `/quizzes`

// @Summary 			Get quizzes
// @Tags					Quiz
// @Success				200 {object} []quizzes.Quiz "Quizzes Out"
// @Produce				json
// @Router				/quiz/quizzes [get]
func (qc *QuizController) GetQuizzes(c *fiber.Ctx, session *users.Session) error {
	quizzesOut, err := qc.quizService.GetQuizzes()
	if err != nil {
		response := helper.NewHTTPResponse(err.Error(), ``)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := helper.NewHTTPResponse(``, quizzesOut)
	return c.Status(fiber.StatusOK).JSON(response)
}

const CreateQuizAction = `/create-quiz`

// @Summary 			Create a quiz
// @Tags					Quiz
// @Param 				requestBody  body  request.CreateQuizIn  true  "Create Quiz In"
// @Success				200 {object} response.CreateQuizOut "Create Quiz Out"
// @Produce				json
// @Router				/quiz/create-quiz [post]
func (qc *QuizController) CreateQuiz(c *fiber.Ctx, session *users.Session) error {
	createQuizIn, err := helper.ReadBody[request.CreateQuizIn](c)
	if err != nil {
		response := helper.NewHTTPResponse(err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if createQuizIn.UserId == 0 {
		createQuizIn.UserId = session.UserID
	}

	createQuizOut, err := qc.quizService.CreateQuiz(createQuizIn)
	if err != nil {
		response := helper.NewHTTPResponse(err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := helper.NewHTTPResponse(``, createQuizOut)
	return c.Status(fiber.StatusOK).JSON(response)
}