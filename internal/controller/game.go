package controller

import (
	"log"
	"triva/internal/service"

	"github.com/gofiber/contrib/websocket"
)

type GameController struct {
	GamePrefix string
	netService *service.NetService
}

func NewGameController(ns *service.NetService) *GameController {
	return &GameController{
		GamePrefix: `/game`,
		netService: ns,
	}
}

const (
	PlayerAction = `/player`
)

func (gc *GameController) Player(conn *websocket.Conn) {
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

		gc.netService.OnIncomingMessage(conn, mt, msg)
	}
}

const (
	HostAction = `/player`
)

func (gc *GameController) Host(conn *websocket.Conn) {
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

		gc.netService.OnIncomingMessage(conn, mt, msg)
	}
}