package biz

import (
	"context"
	"fmt"
	v1 "learn-k8s/api/goods/v1"
)

type Goods struct {
	GoodsID   int64
	GoodsName string
	Price     float32
}

// GetGoods get goods by id
func (uc *GetwayUsecase) GetGoods(ctx context.Context, id int64) (*Goods, error) {
	one, err := uc.goods.GetGoods(ctx, &v1.GetGoodsRequest{Id: id})

	if err != nil {
		return nil, err
	}

	return &Goods{
		GoodsID:   one.Id,
		GoodsName: one.GetName(),
		Price:     one.Price,
	}, err
}

func (uc *GetwayUsecase) ListAllGoods(ctx context.Context) ([]*Goods, error) {
	reply, err := uc.goods.ListGoods(ctx, &v1.ListGoodsRequest{})
	if err != nil {
		fmt.Println("[list]:", err)
		return nil, err
	}

	items := reply.GetList()

	var list = make([]*Goods, len(items))
	for i, item := range items {
		list[i] = &Goods{
			GoodsID:   item.GetId(),
			GoodsName: item.Name,
			Price:     item.Price,
		}
	}
	return list, err
}
