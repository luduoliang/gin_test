package main

import (
	"github.com/gin-gonic/gin"
	"gin_test/router"
	"gin_test/config"
	"gin_test/common"
)

func main() {
	//New不使用中间间，Default使用中间间
	r := gin.New()

	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())

	//初始化SESSION
	common.InitSession(r)
	//视图目录
	r.LoadHTMLGlob("views/**/**/*")
	//静态资源目录
	r.Static("/static", "./static")
	//初始化路由
	router.InitRouter(r)

	r.Run(":" + config.Config.Httpport) // listen and serve on 0.0.0.0:8080
}