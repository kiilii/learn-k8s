package main

import (
	"flag"
	"fmt"

	"learn-k8s/app/greet/greet"
	"learn-k8s/app/greet/internal/config"
	"learn-k8s/app/greet/internal/server"
	"learn-k8s/app/greet/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/gateway"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var gatewayFile = flag.String("g", "etc/gateway.yaml", "the config file")

var configFile = flag.String("f", "etc/greet.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	sg := service.NewServiceGroup()
	defer sg.Stop()

	// grpc service
	sg.Add(zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		greet.RegisterGreetServer(grpcServer, server.NewGreetServer(ctx))

		// if c.Mode == service.DevMode || c.Mode == service.TestMode {
		reflection.Register(grpcServer)
		// }
	}))

	// http service
	var gc gateway.GatewayConf
	conf.MustLoad(*gatewayFile, &gc)
	sg.Add(gateway.MustNewServer(gc))

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	fmt.Printf("Starting http server at %s...\n", fmt.Sprintf("%s:%d", gc.Host, gc.Port))

	sg.Start()
}
