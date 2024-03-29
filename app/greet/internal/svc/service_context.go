package svc

import (
	"learn-k8s/app/greet/internal/config"

	"github.com/zeromicro/go-queue/kq"
)

type ServiceContext struct {
	Config config.Config

	Pusher *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		Pusher: kq.NewPusher(c.LoggerKafkaQueue.Brokers, c.LoggerKafkaQueue.Topic),
	}
}
