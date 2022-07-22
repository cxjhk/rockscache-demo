package main

import (
	"awesome/internal/config"
	"awesome/internal/handler"
	"awesome/internal/svc"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var (
	configFile = flag.String("f", "etc/friend.yaml", "the config file")
)

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	//zero service
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}