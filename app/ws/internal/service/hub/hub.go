package hub

import (
	"context"
	"learn-k8s/app/ws/internal/config"
	"learn-k8s/app/ws/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/threading"
)

type Service struct {
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	svcGroup *service.ServiceGroup
	c        config.Config
	logx.Logger
}

func NewService(svcCtx *svc.ServiceContext, cfg config.Config) *Service {
	m := &Service{
		c:        cfg,
		svcCtx:   svcCtx,
		svcGroup: service.NewServiceGroup(),
		Logger:   logx.WithContext(context.Background()),
	}

	// 统计数据
	// m.svcGroup.Add(kq.MustNewQueue(cfg.AnalyzeGameRoundQueue, xtrace.WithKqHandler(m.GameStatEventHandler))) // 游戏局数
	// m.svcGroup.Add(kq.MustNewQueue(cfg.AppReportEventQueue, xtrace.WithKqHandler(m.AppEventHandler)))
	// m.svcGroup.Add(kq.MustNewQueue(cfg.RegisterEventQueue, xtrace.WithKqHandler(m.RegisterHandler)))
	// m.svcGroup.Add(kq.MustNewQueue(cfg.TagsChangeQueue, xtrace.WithKqHandler(m.NpOrderSuccess)))

	return m
}

func (s *Service) SendMessage(k, v string) error {
	// s.svcCtx.WebsocketHub.
	return nil
}

func (s *Service) Start() {
	threading.GoSafe(func() {
		if s.svcGroup != nil {
			s.svcGroup.Start()
		}
	})

	s.Infof("hub service start")
}

func (s *Service) Stop() {
	if s.svcGroup != nil {
		s.svcGroup.Stop()
	}

	s.Infof("hub service stop")
}
