package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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

func (ns *NetService) OnIncomingMessage(conn *websocket.Conn, mt int, msg []byte) {
	str := string(msg)
	parts := strings.Split(str, ":")
	cmd := parts[0]
	argument := parts[1]

	switch cmd {
	case QUIZ_EVENT_HOST:
		{
			fmt.Println("host quiz:", argument)
			ns.host = conn
			ns.tick = 100
			go func() {
				ns.tick--
				time.Sleep(time.Second)
			}()
			ns.host.WriteMessage(websocket.TextMessage, []byte(strconv.Itoa(ns.tick)))
			break
		}
	case QUIZ_EVENT_JOIN:
		{
			fmt.Println("join code:", argument)
			ns.host.WriteMessage(websocket.TextMessage, []byte("A player joined !!"))
			break
		}
	}
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