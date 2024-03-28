package config

import (
	"learn-k8s/library/xlog"

	"github.com/zeromicro/go-zero/gateway"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	LoggerKafkaQueue xlog.LoggerWriterConfig

	Gateway gateway.GatewayConf
}
