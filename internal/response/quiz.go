package response

import "triva/internal/repository/quizzes"

type CreateQuizOut struct {
	Quiz	*quizzes.Quiz	`json:"quiz" form:"quiz"`
} // @name CreateQuizOut

type CreateQuestionAndChoicesOut struct {
	Quiz	*quizzes.Quiz	`json:"quiz" form:"quiz"`
} // @name CreateQuestionAndChoicesOut