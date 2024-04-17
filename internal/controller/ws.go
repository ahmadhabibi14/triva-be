package controller

import (
	"log"
	"triva/internal/service"

	"github.com/gofiber/contrib/websocket"
)

type WebsocketController struct {
	netService *service.NetService
}

func NewWebsocketController(ns *service.NetService) *WebsocketController {
	return &WebsocketController{netService: ns}
}

func (wsc *WebsocketController) WS(conn *websocket.Conn) {
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