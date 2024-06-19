package quizzes

import (
	"time"
	"triva/internal/bootstrap/database"
)

const TABLE_QuizQuestion string = `quiz_question`

type QuizQuestion struct {
	Db *database.Database `db:"-" json:"-"`

	Id         uint64     `db:"id" json:"id"`
	QuizId     uint64     `db:"quiz_id" json:"quiz_id"`
	Name       string    `db:"name" json:"name"`
	IsUseImage bool      `db:"is_use_image" json:"is_use_image"`
	ImageURL   string    `db:"image_url" json:"image_url"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt  time.Time `db:"deleted_at" json:"deleted_at"`

	Choices []QuizChoice `db:"-" json:"choices"`
} // @name QuizQuestion

func NewQuizQuestionMutator(Db *database.Database) *QuizQuestion {
	return &QuizQuestion{Db: Db}
}

func (qq *QuizQuestion) Insert() error {
	query := `INSERT INTO ` + TABLE_QuizQuestion + `
(quiz_id, name, is_use_image,  image_url, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (id) DO NOTHING
RETURNING id, quiz_id, name, is_use_image,  image_url, created_at, updated_at`

	if err := qq.Db.DB.QueryRowx(query,
		qq.QuizId, qq.Name, qq.IsUseImage, qq.ImageURL, time.Now(), time.Now(),
	).StructScan(qq); err != nil {
		return	err
    }

	return nil
}

func (qq *QuizQuestion) Select(quizId int) ([]QuizQuestion, error) {
    query := `SELECT id, quiz_id, name, is_use_image,  image_url, created_at, updated_at FROM ` + TABLE_QuizQuestion + ` WHERE quiz_id = $1`

    rows, err := qq.Db.DB.Queryx(query, quizId)
    if err != nil {
        return nil, err
    }

    quizQuestions := []QuizQuestion{}
    for rows.Next() {
        quizQuestion := QuizQuestion{}
        if err := rows.StructScan(&quizQuestion); err != nil {
            return nil, err
        }

        quizQuestions = append(quizQuestions, quizQuestion)
    }

    return quizQuestions, nil
}
