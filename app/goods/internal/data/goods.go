package data

import (
	"context"

	"goods/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type goodsRepo struct {
	data *Data
	log  *log.Helper
}

// NewGoodsRepo .
func NewGoodsRepo(data *Data, logger log.Logger) biz.GoodsRepo {
	return &goodsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *goodsRepo) Save(ctx context.Context, g *biz.Goods) (*biz.Goods, error) {
	return g, nil
}

func (r *goodsRepo) Update(ctx context.Context, g *biz.Goods) (*biz.Goods, error) {
	return g, nil
}

func (r *goodsRepo) FindByID(context.Context, int64) (*biz.Goods, error) {
	return nil, nil
}

func (r *goodsRepo) ListByHello(context.Context, string) ([]*biz.Goods, error) {
	return nil, nil
}

func (r *goodsRepo) ListAll(context.Context) ([]*biz.Goods, error) {
	return nil, nil
}
