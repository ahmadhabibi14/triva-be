package service

import (
	"fmt"
	"math"
	"sort"
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
	IntermissionState
	RevealState
	EndState
)

type LeaderboardEntry struct {
	Name string `json:"name"`
	Points int `json:"points"`
}

const TABLE_Game string = `Game` 

type GameService struct {
	Id string `json:"id"`
	Quiz quizzes.Quiz `json:"quiz"`
	CurrentQuestion int `json:"current_question"`
	Code string `json:"code"`
	State GameState `json:"game_state"`
	Ended bool `json:"ended"`
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

func (gs *GameService) StartOrSkip() {
	if gs.State == LobbyState {
		gs.Start()
	} else {
		gs.NextQuestion()
	}
}

func (gs *GameService) Start() {
	gs.ChangeState(PlayState)
	gs.NextQuestion()

	go func ()  {
		defer helper.Recover()

		for {
			if gs.Ended {
				return
			}
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

func (gs *GameService) End() {
	gs.Ended = true
	gs.ChangeState(EndState)
}

func (gs *GameService) NextQuestion() {
	gs.CurrentQuestion++

	if gs.CurrentQuestion >= len(gs.Quiz.Questions) {
		gs.End()
		return
	}

	gs.ResetPlayerAnswerStates()
	gs.ChangeState(PlayState)
	gs.Time = 60

	gs.NetService.SendPacket(gs.Host, QuestionShowPacket{
		Question: gs.getCurrentQuestion(),
	})
}

func (gs *GameService) Reveal() {
	gs.Time = 5
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
				gs.Intermission()
				break
			}
		case IntermissionState:
			{
				gs.NextQuestion()
			}
		}
	}
}

func (gs *GameService) Intermission() {
	gs.Time = 30
	gs.ChangeState(IntermissionState)
	gs.NetService.SendPacket(gs.Host, LeaderboardPacket{
		Points: gs.getLeaderboard(),
	})
}

func (gs *GameService) getLeaderboard() []LeaderboardEntry {
	sort.Slice(gs.Players, func(i, j int) bool {
		return gs.Players[i].Points > gs.Players[j].Points
	})

	leaderboard := []LeaderboardEntry{}
	for i := 0; i < int(math.Min(3, float64(len(gs.Players)))); i++ {
		player := gs.Players[i]
		leaderboard = append(leaderboard, LeaderboardEntry{
			Name: player.Name,
			Points: player.Points,
		})
	}

	return leaderboard
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
	} else {
		player.LastAwardedPoints = 0
	}
	player.Answered = true

	if len(gs.getAnswerPlayers()) == len(gs.Players) {
		gs.Reveal()
	}
}