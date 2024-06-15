package request

type CreateQuizIn struct {
	Name 		string	`json:"name" validate:"required"`
	UserId 	uint64	`json:"user_id"`
} // @name CreateQuizIn

type quizChoice struct {
	Name 			string	`json:"name" validate:"required"`
	IsCorrect bool		`json:"is_correct" validate:"required"`
}

type CreateQuestionAndChoicesIn struct {
	QuizId		uint64				`json:"quiz_id" validate:"required"`
	Question	uint64				`json:"question" validate:"required"`
	Choices		[]quizChoice	`json:"choices" validate:"required"`
} // @name CreateQuestionAndChoicesIn