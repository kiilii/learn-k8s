package data

import (
	"api-getway/internal/conf"

	"github.com/go-kratos/kratos/v2/contrib/registry/etcd"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var ns = "api-getway"
var ProviderSet = wire.NewSet(NewDiscovery, NewRegistrar)

// func NewEtcdConf() *conf.Naming {
// 	return nil
// }

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
	return etcd.New(client)
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

	return etcd.New(client)
}
