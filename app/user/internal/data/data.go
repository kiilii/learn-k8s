package data

import (
	"context"
	"user/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

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

	// db, err := gorm.Open(postgres.New(postgres.Config{
	// 	DriverName:           c.GetDatabase().GetDriver(),
	// 	DSN:                  c.GetDatabase().GetSource(),
	// 	PreferSimpleProtocol: true,
	// }), &gorm.Config{})

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
