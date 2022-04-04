package biz

import (
	"learn-k8s/app/goods/service/internal/conf"

	"github.com/go-kratos/kratos/v2/contrib/registry/etcd"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewGoodsUsecase, NewDiscovery, NewRegistrar)

var ns = "/learn-k8s"

func NewDiscovery(c *conf.Naming) registry.Discovery {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   c.Etcd.GetAddr(),
		Username:    c.Etcd.GetUsername(),
		Password:    c.Etcd.GetPassword(),
		DialTimeout: c.Etcd.Timeout.AsDuration(),
	})

	if err != nil {
		panic(err)
	}
	return etcd.New(client, etcd.Namespace(ns))
}

func NewRegistrar(c *conf.Naming) registry.Registrar {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   c.Etcd.GetAddr(),
		Username:    c.Etcd.GetUsername(),
		Password:    c.Etcd.GetPassword(),
		DialTimeout: c.Etcd.Timeout.AsDuration(),
	})

	if err != nil {
		panic(err)
	}

	return etcd.New(client, etcd.Namespace(ns))
}
