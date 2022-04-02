package client

import (
	"goods/internal/conf"

	"github.com/go-kratos/kratos/v2/contrib/registry/etcd"
	"github.com/go-kratos/kratos/v2/registry"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// namespage
const ns = "goods-service"

func NewEtcdClient(c *conf.Naming) registry.Registrar {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:            c.GetEtcd().GetAddr(),
		Username:             c.Etcd.GetUsername(),
		Password:             c.Etcd.Password,
		DialTimeout:          c.Etcd.Timeout.AsDuration(),
		DialKeepAliveTime:    0,
		DialKeepAliveTimeout: 0,
	})

	if err != nil {
		panic(err)
	}

	return etcd.New(client)
}
