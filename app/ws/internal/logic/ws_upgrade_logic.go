package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"learn-k8s/app/ws/internal/svc"
)

type WsUpgradeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWsUpgradeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WsUpgradeLogic {
	return &WsUpgradeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WsUpgradeLogic) WsUpgrade() error {
	// todo: add your logic here and delete this line

	return nil
}
