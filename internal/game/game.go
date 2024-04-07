package game

import "github.com/jmoiron/sqlx"

type Game struct {
	DB *sqlx.DB `db:"-" json:"-"`

	Id string `db:"id" json:"id"`
	QuizId string `db:"quiz_id" json:"quiz"`
	CurrentQuestion int64 `db:"current_question" json:"current_question"`
	Code string `db:"code" json:"code"`
}