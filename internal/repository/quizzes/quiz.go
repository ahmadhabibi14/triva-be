package quizzes

import (
	"errors"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

const TABLE_Quiz string = `Quiz` 

type Quiz struct {
	DB *sqlx.DB `db:"-" json:"-"`

	Id string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at"`
	Questions []QuizQuestion `db:"-" json:"questions"`
}

func NewQuizMutator(db *sqlx.DB) *Quiz { return &Quiz{DB: db} }

func (q *Quiz) GetQuizzes() (quizzes []Quiz, err error) {
	query := `SELECT
		COALESCE(id, '') id, COALESCE(name, '') name
		FROM ` + TABLE_Quiz + `
		ORDER BY name DESC`
	
	err = q.DB.Select(&quizzes, query)
	
	return
}

func (q *Quiz) FindById(id string) error {
	query := `SELECT id, name FROM ` + TABLE_Quiz + ` WHERE id = $1 LIMIT 1`
	err := q.DB.Get(q, query, strings.TrimSpace(id))
	if err != nil {
		return errors.New(`quiz not found`)
	}

	return nil
}