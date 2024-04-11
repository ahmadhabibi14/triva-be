package games

import (
	"bwizz/helper"
	"bwizz/internal/repository/quizzes"
	"fmt"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/jmoiron/sqlx"
)

type Player struct {
	Name string `json:"name"`
	Connection *websocket.Conn
}

type GameState int

const (
	LobbyState GameState = iota
	PlayState
	RevealState
	EndState
)

const TABLE_Game string = `Game` 

type Game struct {
	DB *sqlx.DB `db:"-" json:"-"`

	Id string `db:"id" json:"id"`
	QuizId string `db:"quiz_id" json:"quiz_id"`
	CurrentQuestion int64 `db:"current_question" json:"current_question"`

	State GameState `db:"-" json:"game_state"`

	Quiz quizzes.Quiz `db:"-" json:"quiz"`
	Code string `db:"code" json:"code"`
	Player []Player

	Host *websocket.Conn `db:"-" json:"-"`
}

func NewGameMutator(db *sqlx.DB, quiz quizzes.Quiz, host *websocket.Conn) *Game {
	return &Game{
		DB: db,
		Quiz: quiz,
		Code: helper.GenerateGameCode(),
		Player: []Player{},
		State: LobbyState,
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

func (g *Game) OnPlayerJoin(name string, conn *websocket.Conn) {
	fmt.Println(name, "joined the game")
	g.Player = append(g.Player, Player{
		Name: name,
		Connection: conn,
	})
}