package quizzes

import (
	"time"
	"triva/internal/bootstrap/database"
)

const TABLE_QuizQuestion string = `QuizQuestion`

type QuizQuestion struct {
	Db *database.Database `db:"-" json:"-"`

	Id 					string 		`db:"id" json:"id"`
	Name 				string 		`db:"name" json:"name"`
	IsUseImage 	bool 			`db:"isUseImage" json:"isUseImage"`
	ImageURL		string		`db:"imageUrl" json:"imageUrl"`
	CreatedAt 	time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt 	time.Time `db:"updatedAt" json:"updatedAt"`
	DeletedAt 	time.Time `db:"deletedAt" json:"deletedAt"`

	Choices []QuizChoice `db:"-" json:"choices"`
}

func NewQuizQuestionMutator(Db *database.Database) *QuizQuestion {
	return &QuizQuestion{Db: Db}
}