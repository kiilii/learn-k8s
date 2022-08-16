package service

import (
	"context"
	v1 "learn-k8s/api/getway/v1"
)

func (s *GetwayService) GetGoods(ctx context.Context, in *v1.GetGoodsRequest) (*v1.GetGoodsReply, error) {
	one, err := s.uc.GetGoods(ctx, in.GetGoodsId())
	if err != nil {
		return nil, err
	}
	return &v1.GetGoodsReply{
		GoodsId:  one.GoodsID,
		GoodName: one.GoodsName,
		Price:    one.Price,
	}, nil
}

func (s *GetwayService) ListGoods(ctx context.Context, in *v1.ListGoodsRequest) (*v1.ListGoodsReply, error) {
	items, err := s.uc.ListAllGoods(ctx)
	if err != nil {
		return nil, err
	}

	// items := []*biz.Goods{
	// 	{GoodsID: 1, GoodsName: "豆浆", Price: 3},
	// }

	var list []*v1.Goods = make([]*v1.Goods, len(items))
	for i, item := range items {
		list[i] = &v1.Goods{
			GoodsId:  item.GoodsID,
			GoodName: item.GoodsName,
			Price:    item.Price,
		}
	}

	return &v1.ListGoodsReply{List: list}, nil
}
