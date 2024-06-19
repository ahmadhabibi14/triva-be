package response

import "triva/internal/repository/quizzes"

type CreateQuizOut struct {
	Quiz	*quizzes.Quiz	`json:"quiz" form:"quiz"`
} // @name CreateQuizOut

type CreateQuestionOut struct {
    Question	*quizzes.QuizQuestion	`json:"question" form:"question"`
} // @name CreateQuestionOut
