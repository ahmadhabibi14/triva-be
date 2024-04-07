package quizzes

import (
	"bwizz/configs"

	"github.com/jmoiron/sqlx"
)

const TABLE_QuizChoice string = `QuizChoice`

type QuizChoice struct {
	DB *sqlx.DB `db:"-" json:"-"`

	Id string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Correct bool `db:"correct" json:"correct"`
}

func NewQuizChoiceMutator() *QuizChoice {
	return &QuizChoice{DB: configs.PostgresDB}
}