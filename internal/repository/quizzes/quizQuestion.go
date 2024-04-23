package quizzes

import (
	"time"

	"github.com/jmoiron/sqlx"
)

const TABLE_QuizQuestion string = `QuizQuestion`

type QuizQuestion struct {
	DB *sqlx.DB `db:"-" json:"-"`

	Id string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at"`

	Choices []QuizChoice `db:"-" json:"choices"`
}

func NewQuizQuestionMutator(db *sqlx.DB) *QuizQuestion { return &QuizQuestion{DB: db} }