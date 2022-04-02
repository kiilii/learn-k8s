package client

import (
	"api-getway/internal/conf"

	"github.com/go-kratos/kratos/v2/contrib/registry/etcd"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// namespace
const ns = "api-getway"

func NewEtcdClient(c *conf.Naming) *etcd.Registry {
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
