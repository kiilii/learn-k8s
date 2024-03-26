package logic

import (
	"context"
	"fmt"

	"learn-k8s/app/greet/greet"
	"learn-k8s/app/greet/internal/svc"

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

func (l *PingLogic) Ping(in *greet.Request) (*greet.Response, error) {
	// todo: add your logic here and delete this line

	return &greet.Response{
		Pong: fmt.Sprintf("greet.Greet/Ping: %s", in.GetPing()),
	}, nil
}
