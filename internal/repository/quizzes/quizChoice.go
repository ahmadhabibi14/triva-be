package quizzes

import (
	"time"
	"triva/internal/bootstrap/database"
)

const TABLE_QuizChoice string = `QuizChoice`

type QuizChoice struct {
	Db *database.Database `db:"-" json:"-"`

	Id 				string 		`db:"id" json:"id"`
	Name 			string 		`db:"name" json:"name"`
	IsCorrect bool 			`db:"isCorrect" json:"isCorrect"`
	CreatedAt time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt" json:"updatedAt"`
	DeletedAt time.Time `db:"deletedAt" json:"deletedAt"`
}

func NewQuizChoiceMutator(Db *database.Database) *QuizChoice {
	return &QuizChoice{Db: Db}
}