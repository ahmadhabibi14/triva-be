package quizzes

import (
	"errors"
	"strings"
	"time"
	"triva/helper"

	"github.com/jmoiron/sqlx"
)

const TABLE_Quiz string = `Quiz` 

type Quiz struct {
	DB *sqlx.DB `db:"-" json:"-"`

	Id string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	UserId string `db:"user_id" json:"user_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at"`
	Questions []QuizQuestion `db:"-" json:"questions"`
}

func NewQuizMutator(db *sqlx.DB) *Quiz { return &Quiz{DB: db} }

func (q *Quiz) GetQuizzes() (quizzes []Quiz, err error) {
	query := `SELECT
		COALESCE(id, '') id, COALESCE(name, '') name
		FROM ` + TABLE_Quiz + ` WHERE user_id = ` + q.UserId + `
		ORDER BY name DESC`
	
	err = q.DB.Select(&quizzes, query)
	
	return
}

func (q *Quiz) FindById(id string) error {
	query := `SELECT * FROM ` + TABLE_Quiz + ` WHERE id = $1 LIMIT 1`
	err := q.DB.Get(q, query, strings.TrimSpace(id))
	if err != nil {
		return errors.New(`quiz not found`)
	}

	return nil
}

func (q *Quiz) Insert() error {
	query := `INSERT INTO ` + TABLE_Quiz + `
(id, name, user_id, created_at) VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (user_id) DO NOTHING
RETURNING id, name, user_id, created_at, updated_at`

	if err := q.DB.QueryRowx(query,
		helper.RandString(35), q.Name, q.UserId, time.Now(),
	).StructScan(q); err != nil { return err }

	return nil
}

func (q *Quiz) UpdateById() error {
	query := `UPDATE ` + TABLE_Quiz + `
SET name = $1, updated_at = $2
WHERE id = $3
RETURNING id, name, user_id, created_at, updated_at`

	if err := q.DB.QueryRowx(query,
		q.Name, time.Now(), q.Id,
	).StructScan(q); err != nil { return err }

	return nil
}