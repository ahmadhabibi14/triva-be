package game

import (
	"bwizz/helper"
	"bwizz/internal/repository/quizzes"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/jmoiron/sqlx"
)

const TABLE_Game string = `Game` 

type Game struct {
	DB *sqlx.DB `db:"-" json:"-"`

	Id string `db:"id" json:"id"`
	QuizId string `db:"quiz_id" json:"quiz_id"`
	CurrentQuestion int64 `db:"current_question" json:"current_question"`
	Quiz quizzes.Quiz `db:"-" json:"quiz"`
	Code string `db:"code" json:"code"`

	Host *websocket.Conn `db:"-" json:"-"`
}

func NewGameMutator(db *sqlx.DB, quiz quizzes.Quiz, host *websocket.Conn) *Game {
	return &Game{
		DB: db,
		Quiz: quiz,
		Code: helper.GenerateGameCode(),
		Host: host,
	}
}

func (g *Game) Start() {
	go func() {
		defer helper.Recover()
		for {
			g.Tick()
			time.Sleep(time.Second)
		}
	}()
}

func (g *Game) Tick() {

}