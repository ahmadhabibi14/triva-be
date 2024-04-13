package service

import (
	"bwizz/internal/repository/quizzes"
	"errors"
	"fmt"
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/contrib/websocket"
	"github.com/jmoiron/sqlx"
)

const (
	QUIZ_EVENT_HOST = `host`
	QUIZ_EVENT_JOIN = `join`
)

type NetService struct {
	db *sqlx.DB
	quizService *QuizService

	games []*GameService
}

func NewNetService(qs *QuizService, db *sqlx.DB) *NetService {
	return &NetService{
		quizService: qs,
		db: db,
		games: []*GameService{},
}
}

type ConnectPacket struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type HostGamePacket struct {
	QuizId string `json:"quizId"`
}

type QuestionShowPacket struct {
	Question quizzes.QuizQuestion `json:"question"`
}

type ChangeGameStatePacket struct {
	State GameState `json:"state"`
}

type PlayerJoinPacket struct {
	Player Player `json:"player"`
}

type StartGamePacket struct {

}

const (
	PACKET_CONNECT uint8 = iota
	PACKET_HOST
	PACKET_QUESTION
	PACKET_STATE
	PACKET_PLAYER_JOIN
	PACKET_START_GAME
)

func (ns *NetService) packetIdToPacket(packetId uint8) any {
	switch packetId {
	case PACKET_CONNECT:
		{
			return &ConnectPacket{}
		}
	case PACKET_HOST:
		{
			return &HostGamePacket{}
		}
	case PACKET_START_GAME:
		{
			return &StartGamePacket{}
		}
	}

	return nil
}

func (ns *NetService) packetToPacketId(packet any) (uint8, error) {
	switch packet.(type) {
	case QuestionShowPacket:
		{
			return PACKET_QUESTION, nil
		}
	case ChangeGameStatePacket:
		{
			return PACKET_STATE, nil
		}
	case PlayerJoinPacket:
		{
			return PACKET_PLAYER_JOIN, nil
		}
	}

	return 0, errors.New("invalid packet type")
}

func (ns *NetService) getGameByCode(code string) *GameService {
	for _, g := range ns.games {
		if g.Code == code {
			return g
		}
	}

	return nil
}

func (ns *NetService) getGameByHost(host *websocket.Conn) *GameService {
	for _, game := range ns.games {
		if game.Host == host {
			return game
		}
	}

	return nil
}

func (ns *NetService) OnIncomingMessage(conn *websocket.Conn, mt int, msg []byte) {
	if len(msg) < 2 {
		log.Println(`message length is less than 2`)
		return
	}

	fmt.Println(`MSG:`, string(msg))
	
	packetId := msg[0]
	data := msg[1:]

	packet := ns.packetIdToPacket(packetId)
	if packet == nil {
		return
	}

	err := json.Unmarshal(data, &packet)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch data := packet.(type) {
	case *ConnectPacket:
		{
			game := ns.getGameByCode(data.Code)
			if game == nil {
				return
			}
			
			game.OnPlayerJoin(data.Name, conn)
			break
		}
	case *HostGamePacket:
		{
			quiz := quizzes.NewQuizMutator(ns.db)
			err := quiz.FindById(data.QuizId)
			if err != nil {
				log.Println(`(ns *NetService) OnIncomingMessage()`, err)
				return
			}

			game := NewGameService(*quiz, conn, ns)
			log.Println(`new game:`, game.Code)
			ns.games = append(ns.games, game)

			ns.SendPacket(conn, ChangeGameStatePacket{
				State: LobbyState,
			})
			break
		}
	case *StartGamePacket:
		{
			game := ns.getGameByHost(conn)
			if game == nil {
				return
			}

			game.Start()
			break
		}
	}
}

func (ns *NetService) SendPacket(conn *websocket.Conn, packet any) error {
	bytes, err := ns.PacketToBytes(packet)
	if err != nil {
		return err
	}

	return conn.WriteMessage(websocket.BinaryMessage, bytes)
}

func (ns *NetService) PacketToBytes(packet any) ([]byte, error) {
	packetId, err := ns.packetToPacketId(packet)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(packet)
	if err != nil {
		return nil, err
	}

	final := append([]byte{packetId}, bytes...)
	return final, nil
}