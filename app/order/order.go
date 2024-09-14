package main

import (
	"flag"
	"fmt"

	"learn-k8s/app/order/internal/config"
	"learn-k8s/app/order/internal/server"
	"learn-k8s/app/order/internal/svc"
	"learn-k8s/app/order/order"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/gateway"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	ctx := svc.NewServiceContext(c)

	logx.Infof("config: %+v", c)

	sg := service.NewServiceGroup()
	defer sg.Stop()

	// grpc service
	sg.Add(zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		order.RegisterOrderServer(grpcServer, server.NewOrderServer(ctx))

		// if c.Mode == service.DevMode || c.Mode == service.TestMode {
		reflection.Register(grpcServer)
		// }
	}))

	// http service
	sg.Add(gateway.MustNewServer(c.Gateway))

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	fmt.Printf("Starting http server at %s...\n", fmt.Sprintf("%s:%d", c.Gateway.Host, c.Gateway.Port))

	sg.Start()
}
