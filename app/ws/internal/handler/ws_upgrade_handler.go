package handler

import (
	"net/http"

	"learn-k8s/app/ws/internal/svc"
)

func WsUpgradeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// l := logic.NewWsUpgradeLogic(r.Context(), svcCtx)

		conn, err := svcCtx.WebsocketHub.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		conn.WriteJSON("hello")
		svcCtx.WebsocketHub.AddConn(conn)

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
