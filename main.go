package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

	"luxuewen.com/config"
	"luxuewen.com/database"
	"zx.ccbaidu.com/middleware"
)

type Gins struct {
	ID int `gorose:"id"`
	Name string `gorose:"name"`
	Age int `gorose:"age"`
}

func (g*Gins) TableName() string {
    return "gins"
}


func main () {// 获取默认配置文件配置
	conf := config.InitConfig("./config", "config", "yaml")
	if debug := conf.GetBool("Debug"); !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	var g = Gins{}
	err := database.DB().Table(&g).Fields("id,name,age").Where("age","=",23).OrderBy("id desc").Select()
	// res,err := database.DB().Query("select * from gins", 1)
	if err != nil {
		panic(err.Error())
	}
    fmt.Println(g)
	
	

	server := gin.Default()
	server.Use(middleware.IPAuthenticate())
	server.GET("/test", func (c *gin.Context) {
		fmt.Println(g)
		c.String(http.StatusOK, "hello gin")
	})
	server.Run()
}