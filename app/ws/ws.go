package main

import (
	"flag"
	"fmt"

	"learn-k8s/app/ws/internal/config"
	"learn-k8s/app/ws/internal/handler"
	"learn-k8s/app/ws/internal/service/hub"
	"learn-k8s/app/ws/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/ws-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	sg := service.NewServiceGroup()
	defer sg.Stop()
	server := rest.MustNewServer(c.RestConf)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	sg.Add(server)
	sg.Add(hub.NewService(ctx, c))

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	sg.Start()
}
