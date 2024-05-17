package controller

import (
	"log"
	"triva/internal/service"

	"github.com/gofiber/contrib/websocket"
)

type GameController struct {
	GameAction string
	netService *service.NetService
}

func NewGameController(ns *service.NetService) *GameController {
	return &GameController{
		GameAction: `/game`,
		netService: ns,
	}
}

func (wsc *GameController) Game(conn *websocket.Conn) {
	var (
		mt int
		msg []byte
		err error
	)

	for {
		if mt, msg, err = conn.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}

		wsc.netService.OnIncomingMessage(conn, mt, msg)
	}
}