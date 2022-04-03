package biz

import (
	goodsv1 "api-getway/api/goods/v1"
	userv1 "api-getway/api/user/v1"

	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewGetwayUsecase, NewGoodsServiceClient, NewUserServiceClient)

// GetwayUsecase is a Getway usecase.
type GetwayUsecase struct {
	goods goodsv1.GoodsClient
	user  userv1.UserClient
	log   *log.Helper
}

// NewGetwayUsecase new a Getway usecase.
func NewGetwayUsecase(gs goodsv1.GoodsClient, us userv1.UserClient, logger log.Logger) *GetwayUsecase {
	return &GetwayUsecase{
		goods: gs,
		user:  us,
		log:   log.NewHelper(logger),
	}
}

func NewGoodsServiceClient(reg registry.Discovery, logger log.Logger) goodsv1.GoodsClient {
	var endpoint = "discovery:////microservices/api-getway"

	rpc, err := grpc.DialInsecure(
		context.TODO(),
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(reg),
		grpc.WithMiddleware(
			recovery.Recovery(),
			logging.Client(logger),
		),
	)
	if err != nil {
		panic(err)
	}

	return goodsv1.NewGoodsClient(rpc)
}

func NewUserServiceClient(reg registry.Discovery, logger log.Logger) userv1.UserClient {
	var endpoint = "discovery:////microservices/user"

	rpc, err := grpc.DialInsecure(
		context.TODO(),
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(reg),
		grpc.WithMiddleware(
			recovery.Recovery(),
			logging.Client(logger),
		),
	)
	if err != nil {
		panic(err)
	}
	return userv1.NewUserClient(rpc)
}
