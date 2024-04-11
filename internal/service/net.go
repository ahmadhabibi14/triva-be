package service

import (
	"bwizz/internal/repository/game"
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

	games []*game.Game
}

func NewNetService(qs *QuizService, db *sqlx.DB) *NetService {
	return &NetService{
		quizService: qs,
		db: db,
		games: []*game.Game{},
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

const (
	PACKET_CONNECT uint8 = iota
	PACKET_HOST
	PACKET_QUESTION
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
	}

	return nil
}

func (ns *NetService) packetToPacketId(packet any) (uint8, error) {
	switch packet.(type) {
	case QuestionShowPacket:
		{
			return PACKET_QUESTION, nil
		}
	}

	return 0, errors.New("invalid packet type")
}

func (ns *NetService) OnIncomingMessage(conn *websocket.Conn, mt int, msg []byte) {
	if len(msg) < 2 {
		log.Println(`message length is less than 2`)
		return
	}
	
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
			fmt.Println(data.Name, "wants to join game ", data.Code)
			break
		}
	case *HostGamePacket:
		{
			log.Println(`Quiz ID:`, data.QuizId)
			quiz := quizzes.NewQuizMutator(ns.db)
			err := quiz.FindById(data.QuizId)
			if err != nil {
				log.Println(`(ns *NetService) OnIncomingMessage()`, err)
				return
			}

			ns.games = append(ns.games, game.NewGameMutator(ns.db, *quiz, conn))
			log.Println(`Games:`, ns.games)
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