package quizzes

import (
	"github.com/jmoiron/sqlx"
)

const TABLE_QuizQuestion string = `QuizQuestion`

type QuizQuestion struct {
	DB *sqlx.DB `db:"-" json:"-"`

	Id string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Choices []QuizChoice `db:"-" json:"choices"`
}

func NewQuizQuestionMutator(db *sqlx.DB) *QuizQuestion { return &QuizQuestion{DB: db} }