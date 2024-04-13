package service

import (
	"bwizz/helper"
	"bwizz/internal/repository/quizzes"
	"fmt"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/google/uuid"
)

type Player struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Connection *websocket.Conn `json:"-"`
}

type GameState int

const (
	LobbyState GameState = iota
	PlayState
	RevealState
	EndState
)

const TABLE_Game string = `Game` 

type GameService struct {
	Id string `json:"id"`
	CurrentQuestion int64 `json:"current_question"`
	State GameState `json:"game_state"`
	Quiz quizzes.Quiz `json:"quiz"`
	Code string `json:"code"`
	Players []Player

	Host *websocket.Conn `json:"-"`
	NetService *NetService `json:"-"`
}

func NewGameService(quiz quizzes.Quiz, host *websocket.Conn, ns *NetService) *GameService {
	return &GameService{
		Quiz: quiz,
		Code: helper.GenerateGameCode(),
		Players: []Player{},
		State: LobbyState,
		Host: host,
		NetService: ns,
	}
}

func (g *GameService) Start() {
	go func() {
		defer helper.Recover()
		for {
			g.Tick()
			time.Sleep(time.Second)
		}
	}()
}

func (g *GameService) Tick() {

}

func (g *GameService) OnPlayerJoin(name string, conn *websocket.Conn) {
	fmt.Println(name, "joined the game")

	player := Player{
		Id: fmt.Sprintf("%v", uuid.New()),
		Name: name,
		Connection: conn,
	}

	g.Players = append(g.Players, player)

	g.NetService.SendPacket(conn, ChangeGameStatePacket{
		State: g.State,
	})

	g.NetService.SendPacket(conn, PlayerJoinPacket{
		Player: player,
	})
}