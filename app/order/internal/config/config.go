package config

import (
	"learn-k8s/library/xlog"

	"github.com/zeromicro/go-zero/gateway"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	// 日志写入
	LoggerKafkaQueue xlog.LoggerWriterConfig

	Gateway gateway.GatewayConf
}
