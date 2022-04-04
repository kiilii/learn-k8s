package data

import (
	"learn-k8s/app/api-getway/service/internal/conf"

	"github.com/go-kratos/kratos/v2/contrib/registry/etcd"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var ns = "/learn-k8s"
var ProviderSet = wire.NewSet(NewDiscovery, NewRegistrar)

func NewDiscovery(c *conf.Naming) registry.Discovery {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:            c.Etcd.GetAddr(),
		Username:             c.Etcd.GetUsername(),
		Password:             c.Etcd.GetPassword(),
		DialTimeout:          c.Etcd.Timeout.AsDuration(),
		DialKeepAliveTime:    0,
		DialKeepAliveTimeout: 0,
	})

	if err != nil {
		panic(err)
	}
	return etcd.New(client, etcd.Namespace(ns))
}

func NewRegistrar(c *conf.Naming) registry.Registrar {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:            c.Etcd.GetAddr(),
		Username:             c.Etcd.GetUsername(),
		Password:             c.Etcd.GetPassword(),
		DialTimeout:          c.Etcd.Timeout.AsDuration(),
		DialKeepAliveTime:    0,
		DialKeepAliveTimeout: 0,
	})

	if err != nil {
		panic(err)
	}

	return etcd.New(client, etcd.Namespace(ns))
}
