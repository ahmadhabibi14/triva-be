package quizzes

import (
	"time"
	"triva/internal/bootstrap/database"
)

const TABLE_QuizQuestion string = `QuizQuestion`

type QuizQuestion struct {
	Db *database.Database `db:"-" json:"-"`

	Id string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at"`

	Choices []QuizChoice `db:"-" json:"choices"`
}

func NewQuizQuestionMutator(Db *database.Database) *QuizQuestion {
	return &QuizQuestion{Db: Db}
}