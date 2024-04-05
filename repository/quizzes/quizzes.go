package quizzes

import (
	"bwizz/configs"

	"github.com/jmoiron/sqlx"
)

type Quiz struct {
	DB *sqlx.DB `json:"-"`

	Id string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func NewQuizMutator() *Quiz {
	return &Quiz{DB: configs.PostgresDB}
}

func (q *Quiz) FindAll() ([]Quiz, error) {
	query := `SELECT id, name FROM quizzes ORDER BY name DESC`
	rows, err := q.DB.Queryx(query)
	if err != nil {
		return []Quiz{}, err
	}

	defer rows.Close()
	
	var quizzes []Quiz
	for rows.Next() {
		var quiz Quiz
		err := rows.StructScan(&quiz)
		if err != nil {
			return []Quiz{}, err
		}

		quizzes = append(quizzes, quiz)
	}

	if err := rows.Err(); err != nil {
		return []Quiz{}, err
	}
	
	return quizzes, nil
}