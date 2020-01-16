package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

	"luxuewen.com/config"
	"zx.ccbaidu.com/middleware"
)

func main () {
	// 获取默认配置文件配置
	conf := config.InitConfig("./config", "config", "yaml")
	if debug := conf.GetBool("Debug"); !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	fmt.Println(conf.GetBool("debug"))
	
	server := gin.Default()
	server.Use(middleware.IPAuthenticate())
	server.GET("/test", func (c *gin.Context) {
		c.String(http.StatusOK, "hello gin")
	})
	server.Run()
}