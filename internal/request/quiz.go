package request

type CreateQuizIn struct {
	Name   string `db:"name" json:"name" validate:"required"`
	UserId uint64 `db:"user_id" json:"user_id"`
} // @name CreateQuizIn

// @name CreateQuestionIn
type CreateQuestionIn struct {
	QuizId  uint64        `db:"quiz_id" json:"quiz_id" validate:"required"`
	Name    string        `db:"name" json:"name" validate:"required"`
	Choices []CreateChoiceIn `json:"choices" validate:"required,dive"`
}

// @name CreateOptionIn
type CreateChoiceIn struct {
	Name     string `db:"name" json:"name" validate:"required"`
	IsCorrect bool   `db:"is_correct" json:"is_correct"`
}

