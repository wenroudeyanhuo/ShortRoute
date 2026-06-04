// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"ShotTener/pkg/base62"
	"flag"
	"fmt"

	"ShotTener/internal/config"
	"ShotTener/internal/handler"
	"ShotTener/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/shortener-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	//fmt.Printf("load conf :%#v\n", c)
	//加载了配置之后进行base62的初始化
	base62.MustInit(c.BaseString)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
