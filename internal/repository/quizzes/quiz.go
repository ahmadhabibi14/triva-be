package quizzes

import (
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"
)

const TABLE_Quiz string = `Quiz` 

type Quiz struct {
	DB *sqlx.DB `db:"-" json:"-"`

	Id string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
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