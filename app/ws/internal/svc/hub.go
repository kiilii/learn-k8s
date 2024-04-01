package svc

import (
	"context"
	"learn-k8s/app/ws/internal/config"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

type WebsocketHub struct {
	clients map[*websocket.Conn]bool
	*websocket.Upgrader

	logx.Logger
}

func NewWebsocketHub(c config.Config) *WebsocketHub {
	return &WebsocketHub{
		clients: make(map[*websocket.Conn]bool),

		Logger: logx.WithContext(context.TODO()).WithFields(),

		Upgrader: &websocket.Upgrader{
			HandshakeTimeout: time.Second * time.Duration(c.WebsocketConfig.HandshakeTimeout),
			ReadBufferSize:   c.WebsocketConfig.ReadBufferSize,
			WriteBufferSize:  c.WebsocketConfig.WriteBufferSize,
		},
	}
}

func (w *WebsocketHub) AddConn(conn *websocket.Conn) {
	w.clients[conn] = true
}

func (w *WebsocketHub) Broadcast(data any) {
	for conn := range w.clients {
		err := conn.WriteJSON(data)
		if err != nil {
			// delete(w.clients, conn)
			conn.Close()
		}
	}
}
