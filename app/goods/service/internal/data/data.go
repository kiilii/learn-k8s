package data

import (
	"learn-k8s/app/goods/service/internal/biz"
	"learn-k8s/app/goods/service/internal/conf"
	"sync"

	etcd "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGoodsRepo, NewDiscovery, NewRegistrar)

// Data .
type Data struct {
	// TODO wrapped database client
	// KV *redis.Conn
	m     map[int64]*biz.Goods
	count int64
	mu    *sync.RWMutex
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	return faker(), cleanup, nil
}

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
