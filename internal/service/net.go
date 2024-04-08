package service

import (
	"fmt"

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

func (ns *NetService) OnIncomingMessage(conn *websocket.Conn, mt int, msg []byte) {
	if len(msg) < 1 {
		return
	}
	
	packetId := msg[0]
	data := msg[1:]

	var packet ConnectPacket
	err := json.Unmarshal(data, &packet)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(packetId)
	fmt.Println(packet)
}

func (ns *NetService) SendPacket(conn *websocket.Conn, packet any) error {
	bytes, err := ns.PacketToBytes(packet)
	if err != nil {
		return err
	}

	return conn.WriteMessage(websocket.BinaryMessage, bytes)
}

func (ns *NetService) PacketToBytes(packet any) ([]byte, error) {
	var packetId uint8 = 0

	bytes, err := json.Marshal(packet)
	if err != nil {
		return nil, err
	}

	final := append([]byte{packetId}, bytes...)

	return final, nil
}