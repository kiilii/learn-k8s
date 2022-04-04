package service

import (
	v1 "learn-k8s/api/goods/v1"
	"learn-k8s/app/goods/service/internal/biz"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGoodsService)

// GreeterService is a greeter service.
type GoodsService struct {
	v1.UnimplementedGoodsServer

	uc *biz.GoodsUsecase
}

// NewGreeterService new a greeter service.
func NewGoodsService(uc *biz.GoodsUsecase) *GoodsService {
	return &GoodsService{uc: uc}
}
