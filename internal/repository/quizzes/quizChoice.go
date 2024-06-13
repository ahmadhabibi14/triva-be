package quizzes

import (
	"errors"
	"time"
	"triva/internal/bootstrap/database"
)

const TABLE_QuizChoice string = `quiz_choice`

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

func (qc *QuizChoice) InsertMany(choicesIn []QuizChoice) (choicesOut *[]QuizChoice, err error) {
	if len(choicesIn) == 0 {
		err = errors.New(`quiz choices cannot be empty`)
		return
	}

	query := `INSERT INTO ` + TABLE_QuizChoice + `
(question_id, name, is_correct, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (question_id) DO NOTHING
RETURNING id, question_id, name, is_correct, created_at, updated_at`

	for _, c := range choicesIn {
		err = qc.Db.DB.QueryRowx(query,
			c.QuestionId, c.Name, c.IsCorrect, time.Now(), time.Now(),
		).StructScan(&c)

		if err != nil {
			err = errors.New(`failed insert quiz choices`)
			return
		}
	}

	choicesOut = &choicesIn
	return
}