package quizzes

import (
	"errors"
	"time"
	"triva/internal/bootstrap/database"
)

const TABLE_QuizQuestion string = `quiz_question`

type QuizQuestion struct {
	Db *database.Database `db:"-" json:"-"`

	Id 					string 		`db:"id" json:"id"`
	QuizId 			string 		`db:"quiz_id" json:"quiz_id"`
	Name 				string 		`db:"name" json:"name"`
	IsUseImage 	bool 			`db:"is_use_image" json:"is_use_image"`
	ImageURL		string		`db:"image_url" json:"image_url"`
	CreatedAt 	time.Time `db:"created_at" json:"created_at"`
	UpdatedAt 	time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt 	time.Time `db:"deleted_at" json:"deleted_at"`

	Choices []QuizChoice `db:"-" json:"choices"`
} // @name QuizQuestion

func NewQuizQuestionMutator(Db *database.Database) *QuizQuestion {
	return &QuizQuestion{Db: Db}
}

func (qq *QuizQuestion) Insert() error {
	query := `INSERT INTO ` + TABLE_QuizQuestion + `
(quiz_id, name, is_use_image,  image_url, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (quiz_id) DO NOTHING
RETURNING id, quiz_id, name, is_use_image,  image_url, created_at, updated_at`

	if err := qq.Db.DB.QueryRowx(query,
		qq.QuizId, qq.Name, qq.IsUseImage, qq.ImageURL, time.Now(), time.Now(),
	).StructScan(qq); err != nil {
		return errors.New(`failed to insert a new question`)
	}

	return nil
}