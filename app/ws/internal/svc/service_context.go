package svc

import (
	"learn-k8s/app/ws/internal/config"
)

type ServiceContext struct {
	Config config.Config

	WebsocketHub *WebsocketHub
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		WebsocketHub: NewWebsocketHub(c),
	}
}
