package client

import (
	"goods/internal/conf"

	"github.com/go-kratos/kratos/v2/contrib/registry/etcd"
	"github.com/go-kratos/kratos/v2/registry"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func NewEtcdClient(c *conf.Naming) registry.Registrar {
	client, _ := clientv3.New(clientv3.Config{
		Endpoints:            c.GetEtcd().GetAddr(),
		Username:             c.Etcd.GetUsername(),
		Password:             c.Etcd.Password,
		DialTimeout:          c.Etcd.Timeout.AsDuration(),
		DialKeepAliveTime:    0,
		DialKeepAliveTimeout: 0,
	})

	return etcd.New(client)
}
