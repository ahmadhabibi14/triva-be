package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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

func (c *NetService) OnIncomingMessage(conn *websocket.Conn, mt int, msg []byte) {
	str := string(msg)
	parts := strings.Split(str, ":")
	cmd := parts[0]
	argument := parts[1]

	switch cmd {
	case QUIZ_EVENT_HOST:
		{
			fmt.Println("host quiz:", argument)
			c.host = conn
			c.tick = 100
			go func() {
				c.tick--
				time.Sleep(time.Second)
			}()
			c.host.WriteMessage(websocket.TextMessage, []byte(strconv.Itoa(c.tick)))
			break
		}
	case QUIZ_EVENT_JOIN:
		{
			fmt.Println("join code:", argument)
			c.host.WriteMessage(websocket.TextMessage, []byte("A player joined !!"))
			break
		}
	}
}