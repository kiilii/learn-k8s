package svc

import (
	"learn-k8s/app/ws/internal/config"
)

type ServiceContext struct {
	Config config.Config

	WebsocketConfig struct {
		// HandshakeTimeout specifies the duration for the handshake to complete.
		HandshakeTimeout int64

		// ReadBufferSize and WriteBufferSize specify I/O buffer sizes in bytes. If a buffer
		// size is zero, then buffers allocated by the HTTP server are used. The
		// I/O buffer sizes do not limit the size of the messages that can be sent
		// or received.
		ReadBufferSize, WriteBufferSize int
	}

	WebsocketHub *WebsocketHub
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		WebsocketHub: NewWebsocketHub(),
	}
}
