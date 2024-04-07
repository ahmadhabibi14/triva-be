package quizzes

import (
	"bwizz/configs"

	"github.com/jmoiron/sqlx"
)

const TABLE_QuizQuestion string = `QuizQuestion`

type QuizQuestion struct {
	DB *sqlx.DB `db:"-" json:"-"`

	Id string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Choices *[]QuizChoice `db:"-" json:"choices"`
}

func NewQuizQuestionMutator() *QuizQuestion {
	return &QuizQuestion{DB: configs.PostgresDB}
}