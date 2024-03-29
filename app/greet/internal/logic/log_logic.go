package logic

import (
	"context"

	"learn-k8s/app/greet/greet"
	"learn-k8s/app/greet/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogLogic {
	return &LogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogLogic) Log(in *greet.LogReq) (out *greet.LogRsp, err error) {
	err = l.svcCtx.Pusher.Push(in.GetMsg())
	if err != nil {
		return nil, err
	}

	return &greet.LogRsp{}, nil
}
