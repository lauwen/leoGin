package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
    "github.com/xormplus/xorm"

	"luxuewen.com/config"
	"luxuewen.com/database"
	"zx.ccbaidu.com/middleware"
)

func main () {// 获取默认配置文件配置
	conf := config.InitConfig("./config", "config", "yaml")
	if debug := conf.GetBool("Debug"); !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	res,err := database.DB().Query("select * from gin where id>1 limit 2", 1)
    fmt.Println(res)
	
	

	server := gin.Default()
	server.Use(middleware.IPAuthenticate())
	server.GET("/test", func (c *gin.Context) {
		c.String(http.StatusOK, "hello gin")
	})
	server.Run()
}