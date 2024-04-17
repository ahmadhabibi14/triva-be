package service

import (
	"fmt"
	"math"
	"time"
	"triva/helper"
	"triva/internal/repository/quizzes"

	"github.com/gofiber/contrib/websocket"
	"github.com/google/uuid"
)

type Player struct {
	Id 								string					`json:"id"`
	Name							string					`json:"name"`
	Connection				*websocket.Conn	`json:"-"`
	Points						int							`json:"-"`
	LastAwardedPoints	int							`json:"-"`
	Answered 					bool						`json:"-"`
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
	Quiz quizzes.Quiz `json:"quiz"`
	CurrentQuestion int `json:"current_question"`
	Code string `json:"code"`
	State GameState `json:"game_state"`
	Time int `json:"time"`
	Players []*Player

	Host *websocket.Conn `json:"-"`
	NetService *NetService `json:"-"`
}

func NewGameService(quiz quizzes.Quiz, host *websocket.Conn, ns *NetService) *GameService {
	return &GameService{
		Id: uuid.New().String(),
		Quiz: quiz,
		CurrentQuestion: -1,
		Code: helper.GenerateGameCode(),
		Time: 60,
		Players: []*Player{},
		State: LobbyState,
		Host: host,
		NetService: ns,
	}
}

func (gs *GameService) Start() {
	gs.ChangeState(PlayState)
	gs.NextQuestion()

	go func ()  {
		defer helper.Recover()

		for {
			gs.Tick()
			time.Sleep(time.Second)
		}
	}()
}

func (gs *GameService) ResetPlayerAnswerStates() {
	for _, player := range gs.Players {
		player.Answered = false
	}
}

func (gs *GameService) NextQuestion() {
	gs.CurrentQuestion++

	gs.ResetPlayerAnswerStates()
	gs.ChangeState(PlayState)
	gs.Time = 60

	gs.NetService.SendPacket(gs.Host, QuestionShowPacket{
		Question: gs.getCurrentQuestion(),
	})
}

func (gs *GameService) Reveal() {
	gs.Time = 10
	for _, player := range gs.Players {
		gs.NetService.SendPacket(player.Connection, PlayerRevealPacket{
			Points: player.LastAwardedPoints,
		})
	}

	gs.ChangeState(RevealState)
}

func (gs *GameService) Tick() {
	gs.Time--
	gs.NetService.SendPacket(gs.Host, TickPacket{
		Tick: gs.Time,
	})

	if gs.Time == 0 {
		switch gs.State {
		case PlayState:
			{
				gs.ChangeState(RevealState)
				break
			}
		case RevealState:
			{
				gs.NextQuestion()
				break
			}
		}
	}
}

func (gs *GameService) ChangeState(state GameState) {
	gs.State = state
	gs.BroadcastPacket(ChangeGameStatePacket{
		State: state,
	}, true)
}

func (gs *GameService) BroadcastPacket(packet any, includeHost bool) error {
	for _, player := range gs.Players {
		err := gs.NetService.SendPacket(player.Connection, packet)
		if err != nil {
			return err
		}
	}

	if includeHost {
		err := gs.NetService.SendPacket(gs.Host, packet)
		if err != nil {
			return err
		}
	}

	return nil
}

func (gs *GameService) OnPlayerJoin(name string, conn *websocket.Conn) {
	fmt.Println(name, "joined the game")

	player := Player{
		Id: uuid.New().String(),
		Name: name,
		Connection: conn,
	}

	gs.Players = append(gs.Players, &Player{})

	fmt.Println(`Players:`, gs.Players)

	gs.NetService.SendPacket(conn, ChangeGameStatePacket{
		State: gs.State,
	})

	gs.NetService.SendPacket(gs.Host, PlayerJoinPacket{
		Player: player,
	})
}

func (gs *GameService) getAnswerPlayers() []*Player {
	players := []*Player{}

	for _, player := range gs.Players {
		if player.Answered {
			players = append(players, player)
		}
	}

	return players
}

func (gs *GameService) getCurrentQuestion() quizzes.QuizQuestion {
	return gs.Quiz.Questions[gs.CurrentQuestion]
}

func (gs *GameService) isCorrectChoice(choiceIndex int) bool {
	choices := gs.getCurrentQuestion().Choices
	if choiceIndex < 0 || choiceIndex >= len(choices) {
		return false
	}

	return choices[choiceIndex].Correct
}

func (gs *GameService) getPointsReward() int {
	answered := len(gs.getAnswerPlayers())
	orderReward := 5000 - (1000 * math.Min(4, float64(answered)))
	timeReward := gs.Time * (1000 / 60)

	return int(orderReward) + timeReward
}

func (gs *GameService) OnPlayerAnswer(choice int, player *Player) {
	if gs.isCorrectChoice(choice) {
		player.LastAwardedPoints = gs.getPointsReward()
		player.Points += player.LastAwardedPoints
	}
	player.Answered = true

	if len(gs.getAnswerPlayers()) == len(gs.Players) {
		gs.Reveal()
	}
}