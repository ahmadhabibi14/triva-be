package response

import "triva/internal/repository/quizzes"

type CreateQuizOut struct {
	Quiz	*quizzes.Quiz	`json:"quiz" form:"quiz"`
} // @name CreateQuizOut