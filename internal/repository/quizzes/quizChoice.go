package quizzes

import (
	"time"
	"triva/internal/bootstrap/database"
)

const TABLE_QuizChoice string = `QuizChoice`

type QuizChoice struct {
	Db *database.Database `db:"-" json:"-"`

	Id 					string 		`db:"id" json:"id"`
	QuestionId	string 		`db:"question_id" json:"question_id"`
	Name 				string 		`db:"name" json:"name"`
	IsCorrect 	bool 			`db:"is_correct" json:"is_correct"`
	CreatedAt 	time.Time `db:"created_at" json:"created_at"`
	UpdatedAt 	time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt 	time.Time `db:"deleted_at" json:"deleted_at"`
}

func NewQuizChoiceMutator(Db *database.Database) *QuizChoice {
	return &QuizChoice{Db: Db}
}