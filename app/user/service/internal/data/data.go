package data

import (
	"context"
	"learn-k8s/app/user/service/internal/conf"

	"github.com/go-kratos/kratos/v2/contrib/registry/etcd"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewDiscovery, NewRegistrar)

// Data .
type Data struct {
	// TODO wrapped database client
	// db *gorm.DB
	kv *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	rc := c.GetRedis()
	kv := redis.NewClient(&redis.Options{
		Network: rc.GetNetwork(),
		Addr:    rc.GetAddr(),
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			return cn.Ping(ctx).Err()
		},
		ReadTimeout:  rc.GetReadTimeout().AsDuration(),
		WriteTimeout: rc.WriteTimeout.AsDuration(),
	})

	err := kv.Ping(kv.Context()).Err()
	if err != nil {
		return nil, nil, err
	}

	return &Data{kv: kv}, cleanup, nil
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
