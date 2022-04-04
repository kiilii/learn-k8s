package data

import (
	"learn-k8s/app/goods/service/internal/biz"
	"learn-k8s/app/goods/service/internal/conf"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGoodsRepo)

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
