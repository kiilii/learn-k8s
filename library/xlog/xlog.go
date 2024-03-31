package xlog

import (
	"io"
	"time"

	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/zeromicro/go-queue/kq"
)

var _ (io.Writer) = (*KafkaLoggerWriter)(nil)

type LoggerKafkaWriterConfig struct {
	Brokers       []string `json:",omitempty"`
	Topic         string   `json:",omitempty"`
	FlushInterval int64    `json:",optional"`
	ChunkSize     int      `json:",optional"`
	SASL          plain.Mechanism
}

type KafkaLoggerWriter struct {
	c *LoggerKafkaWriterConfig

	pusher *kq.Pusher
}

// Write implements io.Writer.
func (k *KafkaLoggerWriter) Write(p []byte) (n int, err error) {
	err = k.pusher.Push(string(p))
	if err != nil {
		return
	}
	return len(p), err
}

func NewLoggerWriter(c *LoggerKafkaWriterConfig) *KafkaLoggerWriter {
	var opts []kq.PushOption

	if len(c.Brokers) == 0 {
		panic("brokers must be setting")
	}
	if c.Topic == "" {
		panic("topic must be not empty")
	}

	if c.FlushInterval > 0 {
		opts = append(opts, kq.WithFlushInterval(time.Second*time.Duration(c.FlushInterval)))
	}
	if c.ChunkSize > 0 {
		opts = append(opts, kq.WithChunkSize(c.ChunkSize))
	}

	return &KafkaLoggerWriter{
		c:      c,
		pusher: kq.NewPusher(c.Brokers, c.Topic, opts...),
	}
}
