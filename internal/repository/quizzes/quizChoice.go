package quizzes

import (
	"time"
	"triva/internal/database"
)

const TABLE_QuizChoice string = `QuizChoice`

type QuizChoice struct {
	Db *database.Database `db:"-" json:"-"`

	Id string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Correct bool `db:"correct" json:"correct"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at"`
}

func NewQuizChoiceMutator(Db *database.Database) *QuizChoice {
	return &QuizChoice{Db: Db}
}