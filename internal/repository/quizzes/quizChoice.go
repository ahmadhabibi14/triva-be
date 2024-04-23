package quizzes

import (
	"time"

	"github.com/jmoiron/sqlx"
)

const TABLE_QuizChoice string = `QuizChoice`

type QuizChoice struct {
	DB *sqlx.DB `db:"-" json:"-"`

	Id string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Correct bool `db:"correct" json:"correct"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at"`
}

func NewQuizChoiceMutator(db *sqlx.DB) *QuizChoice { return &QuizChoice{DB: db} }