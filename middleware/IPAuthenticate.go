package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"luxuewen.com/config"
	"luxuewen.com/message"
)

//  通过用户的传值以及配置文件中获取IP黑名单，过滤无权访问IP
func IPAuthenticate (ipList ...string) gin.HandlerFunc {
	return func (c *gin.Context) {
		//  Get the IP blacklist from configure
		ipConfigList := config.InitConfig("./config", "config", "yaml").GetStringSlice("IPBlacklist")

		// 合并配置文件中ip和传入ip
		ipBlacklist := append(ipConfigList, ipList...)

		// 获取客户端ip
		clientIP := c.ClientIP()

		// 验证客户端ip是否在黑名单
		for _, host := range ipBlacklist {
			if clientIP == host {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": http.StatusUnauthorized,
					"message": message.StatusUnauthorized,
				})
				c.Abort()
				break
			}
		}
	}
}