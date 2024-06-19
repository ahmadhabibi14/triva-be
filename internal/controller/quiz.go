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
	createQuizIn, err := helper.ReadBody[request.CreateQuizIn](c, c.Body())
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

const CreateQuestionAction = `/create-question`

// @Summary 			Create question with quiz id
// @Tags				Quiz
// @Param 				requestBody  body  request.CreateQuestionIn  true  "Create Question In"
// @Success				200 {object} response.CreateQuestionOut "Create Question Out"
// @Produce				json
// @Router				/quiz/create-question [post]
func (qc *QuizController) CreateQuestion(c *fiber.Ctx, session *users.Session) error {
	createQuesitonIn, err := helper.ReadBody[request.CreateQuestionIn](c, c.Body())
	if err != nil {
		response := helper.NewHTTPResponse(err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

    //BUG: Check if the quiz id belongs to the user
	// start a transaction before creating a question

	createQuestionOut, err := qc.quizService.CreateQuestion(createQuesitonIn)
	if err != nil {
		response := helper.NewHTTPResponse(err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := helper.NewHTTPResponse(``, createQuestionOut)
	return c.Status(fiber.StatusOK).JSON(response)
}

const GetQuestionAction = `/questions`
// @Summary 			Create question with quiz id
// @Tags				Quiz
// @Param 				requestBody  body  request.CreateQuestionIn  true  "Create Question In"
// @Success				200 {object} response.CreateQuestionOut "Create Question Out"
// @Produce				json
// @Router				/quiz/create-question [post]
func (qc *QuizController) GetQuestions(c *fiber.Ctx, session *users.Session) error {
    //BUG: add a check to see if the quiz belongs to the user

    qid := c.QueryInt("quiz_id", -1)

	if qid == -1 {
		response := helper.NewHTTPResponse("no quiz_id found", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	questions, err := qc.quizService.GetQuestions(qid)
	if err != nil {
		response := helper.NewHTTPResponse(err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := helper.NewHTTPResponse(``, questions)
	return c.Status(fiber.StatusOK).JSON(response)
}

