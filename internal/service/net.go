package service

import (
	"bwizz/internal/repository/quizzes"
	"errors"
	"fmt"
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/contrib/websocket"
)

const (
	QUIZ_EVENT_HOST = `host`
	QUIZ_EVENT_JOIN = `join`
)

type NetService struct {
	quizService *QuizService

	host *websocket.Conn
	tick int
}

func NewNetService(qs *QuizService) *NetService {
	return &NetService{quizService: qs}
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

	switch packet := packet.(type) {
	case *ConnectPacket:
		{
			fmt.Println(packet.Name, "wants to join game ", packet.Code)
			break
		}
	case *HostGamePacket:
		{
			fmt.Println("User wants to host quiz ", packet.QuizId)
			go func() {
				ns.SendPacket(conn, QuestionShowPacket{
					Question: quizzes.QuizQuestion{
						Name: "What is 2+2 ?",
						Choices: []quizzes.QuizChoice{
							{
								Name: "4",
								Correct: true,
							}, {
								Name: "9",
							}, {
								Name: "11",
							}, {
								Name: "Elephant",
							},
						},
					},
				})
			}()
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