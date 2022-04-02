package biz

import (
	goodsv1 "api-getway/api/goods/v1"
	userv1 "api-getway/api/user/v1"
	"fmt"

	"context"

	"github.com/go-kratos/kratos/v2/contrib/registry/etcd"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewGetwayUsecase)

// GetwayUsecase is a Getway usecase.
type GetwayUsecase struct {
	goods goodsv1.GoodsClient
	user  userv1.UserClient
	log   *log.Helper
}

// NewGetwayUsecase new a Getway usecase.
func NewGetwayUsecase(reg *etcd.Registry, logger log.Logger) *GetwayUsecase {
	var ctx = context.TODO()

	rpc, err := grpc.DialInsecure(ctx, grpc.WithDiscovery(reg))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	goods := goodsv1.NewGoodsClient(rpc)
	user := userv1.NewUserClient(rpc)

	var uc = &GetwayUsecase{
		goods: goods,
		user:  user,
		log:   log.NewHelper(logger),
	}

	fmt.Println(uc.goods)
	fmt.Println(uc.user)
	fmt.Println(rpc.GetMethodConfig("/goods/GetGoods"))
	return uc
}
