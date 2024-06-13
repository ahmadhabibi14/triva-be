package quizzes

import (
	"errors"
	"time"
	"triva/internal/bootstrap/database"

	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/S"
)

const TABLE_Quiz string = `quiz` 

type Quiz struct {
	Db *database.Database `db:"-" json:"-"`

	Id 				uint64 					`db:"id" json:"id,omitempty"`
	Name 			string 					`db:"name" json:"name,omitempty"`
	UserId 		uint64					`db:"user_id" json:"user_id,omitempty"`
	CreatedAt time.Time 			`db:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time 			`db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt time.Time 			`db:"deleted_at" json:"deleted_at,omitempty"`
	Questions []QuizQuestion	`db:"-" json:"questions,omitempty"`
} // @name Quiz

func NewQuizMutator(Db *database.Database) *Quiz {
	return &Quiz{Db: Db}
}

func (q *Quiz) GetQuizzes() (quizzes []Quiz, err error) {
	query := `SELECT
		COALESCE(id, '') id, COALESCE(name, '') name
		FROM ` + TABLE_Quiz + ` WHERE userId = ` +  I.UToS(q.UserId) + `
		ORDER BY name DESC`
	
	err = q.Db.DB.Select(&quizzes, query)
	if err != nil {
		err = errors.New(`failed to get quizzes`)
	}

	return
}

func (q *Quiz) FindById(id string) error {
	query := `SELECT * FROM ` + TABLE_Quiz + ` WHERE id = $1 LIMIT 1`
	err := q.Db.DB.Get(q, query, S.Trim(id))
	if err != nil {
		return errors.New(`quiz not found`)
	}

	return nil
}

func (q *Quiz) Insert() error {
	query := `INSERT INTO ` + TABLE_Quiz + `
(name, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4)
RETURNING id, name, user_id, created_at, updated_at`

	if err := q.Db.DB.QueryRowx(query,
		q.Name, q.UserId, time.Now(), time.Now(),
	).StructScan(q); err != nil {
		return errors.New(`failed to insert a new quiz`)
	}

	return nil
}

func (q *Quiz) UpdateNameById() error {
	query := `UPDATE ` + TABLE_Quiz + `
SET name = $1, updated_at = $2
WHERE id = $3
RETURNING id, name, user_id, created_at, updated_at`

	if err := q.Db.DB.QueryRowx(query,
		q.Name, time.Now(), q.Id,
	).StructScan(q); err != nil {
		return errors.New(`failed to update quiz`)
	}

	return nil
}