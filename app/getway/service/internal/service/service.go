package service

import (
	v1 "learn-k8s/api/getway/v1"
	"learn-k8s/app/getway/service/internal/biz"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGetwayService)

// GetwayService is a getway service.
type GetwayService struct {
	v1.UnimplementedGetwayServer

	uc *biz.GetwayUsecase
}

// NewGetwayService new a getway service.
func NewGetwayService(uc *biz.GetwayUsecase) *GetwayService {
	return &GetwayService{uc: uc}
}
