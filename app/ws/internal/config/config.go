package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	WebsocketConfig struct {
		// HandshakeTimeout specifies the duration for the handshake to complete.
		HandshakeTimeout int64

		// ReadBufferSize and WriteBufferSize specify I/O buffer sizes in bytes. If a buffer
		// size is zero, then buffers allocated by the HTTP server are used. The
		// I/O buffer sizes do not limit the size of the messages that can be sent
		// or received.
		ReadBufferSize, WriteBufferSize int
	}
}
