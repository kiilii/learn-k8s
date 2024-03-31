package handler

import (
	"net/http"
	"time"

	"learn-k8s/app/ws/internal/svc"

	"github.com/gorilla/websocket"
)

func WsUpgradeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// l := logic.NewWsUpgradeLogic(r.Context(), svcCtx)

		c := svcCtx.WebsocketConfig
		wsc := websocket.Upgrader{
			HandshakeTimeout: time.Second * time.Duration(c.HandshakeTimeout),
			ReadBufferSize:   c.ReadBufferSize,
			WriteBufferSize:  c.WriteBufferSize,
		}
		conn, err := wsc.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		conn.WriteJSON("hello")

		// if websocket.IsWebSocketUpgrade(r) {

		// } else {
		// 	err := l.WsUpgrade()
		// 	if err != nil {
		// 		httpx.ErrorCtx(r.Context(), w, err)
		// 	} else {
		// 		httpx.Ok(w)
		// 	}
		// }
	}
}
