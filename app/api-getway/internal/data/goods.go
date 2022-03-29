package data

import (
	"api-getway/internal/biz"
	"context"
)

func (r *getwayRepo) FindGoodsByID(ctx context.Context, id int64) (*biz.Goods, error) {
	one := r.data.GetGoods(id)
	if one != nil {
		return one, nil
	}
	return nil, biz.ErrUserNotFound
}

func (r *getwayRepo) ListAllGoods(context.Context) ([]*biz.Goods, error) {
	return r.data.AllGoods(), nil
}
