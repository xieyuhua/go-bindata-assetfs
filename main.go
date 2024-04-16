package main

import (
	"catchace"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var port = flag.String("port", "80", "服务监听端口")

func main() {
	flag.Parse()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// 路由配置
	catchace.RouteConfig(r)
	log.Printf(">> 服务端口:%s", *port)
	err := r.Run(":" + *port)
	log.Fatal(err)
}
