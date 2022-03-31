package service

import (
	"context"

	v1 "goods/api/goods/v1"
)

// SayHello implements helloworld.GreeterServer.
func (s *GoodsService) GetGoods(ctx context.Context, in *v1.GetGoodsRequest) (*v1.GetGoodsReply, error) {
	// g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	// if err != nil {
	// 	return nil, err
	// }
	return &v1.GetGoodsReply{}, nil
}
