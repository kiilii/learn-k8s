package logic

import (
	"context"

	"learn-k8s/app/order/internal/svc"
	"learn-k8s/app/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *order.Request) (*order.Response, error) {
	// todo: add your logic here and delete this line

	return &order.Response{
		Pong: "order.Order/Ping",
	}, nil
}
