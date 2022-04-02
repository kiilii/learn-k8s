package service

import (
	"context"

	v1 "goods/api/goods/v1"
	"goods/internal/biz"
)

// SayHello implements goods.GreeterServer.
func (s *GoodsService) GetGoods(ctx context.Context, in *v1.GetGoodsRequest) (*v1.GetGoodsReply, error) {
	one, err := s.uc.GetGoods(ctx, in.Id)

	return &v1.GetGoodsReply{
		Id:    one.ID,
		Name:  one.Name,
		Price: one.Price,
	}, err
}

func (s *GoodsService) CreateGoods(ctx context.Context, in *v1.CreateGoodsRequest) (*v1.CreateGoodsReply, error) {
	item, err := s.uc.CreateGoods(ctx, &biz.Goods{Name: in.GetName(), Price: in.Price})
	if err != nil {
		return nil, err
	}
	return &v1.CreateGoodsReply{
		Id:    item.ID,
		Name:  item.Name,
		Price: item.Price,
	}, err
}

func (s *GoodsService) ListGoods(ctx context.Context, in *v1.ListGoodsRequest) (*v1.ListGoodsReply, error) {
	items, err := s.uc.ListAll(ctx)

	var list = make([]*v1.Item, len(items))

	for i, item := range items {
		list[i] = &v1.Item{
			Id:    item.ID,
			Name:  item.Name,
			Price: item.Price,
		}
	}

	return &v1.ListGoodsReply{List: list}, err
}
