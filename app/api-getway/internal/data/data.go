package data

import (
	"api-getway/internal/biz"
	"api-getway/internal/conf"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGetwayRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	m *sync.Map
}

type goods struct {
	GoodsID   int64
	GoodsName string
	Price     float64
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	var index int64 = 1
	m := new(sync.Map)
	xx := []*biz.Goods{
		{GoodsID: index, GoodsName: "稀饭", Price: 2.00},
		{GoodsID: index, GoodsName: "卤蛋", Price: 2.00},
		{GoodsID: index, GoodsName: "肉包", Price: 2.00},
		{GoodsID: index, GoodsName: "豆浆", Price: 3.00},
		{GoodsID: index, GoodsName: "杂粮煎饼", Price: 5.00},
	}

	for i, v := range xx {
		m.Store(int64(i), v)
	}

	return &Data{m: m}, cleanup, nil
}

func (data *Data) GetGoods(id int64) *biz.Goods {
	if one, has := data.m.Load(id); has {
		return one.(*biz.Goods)
	}
	return nil
}

func (data *Data) AllGoods() []*biz.Goods {
	var list = make([]*biz.Goods, 0)
	var tmp *biz.Goods

	data.m.Range(func(key, value interface{}) bool {
		tmp = value.(*biz.Goods)
		list = append(list, tmp)
		tmp = nil
		return true
	})

	return list
}
