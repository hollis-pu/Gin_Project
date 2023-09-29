package router

import (
	"Gin_Project/src/18_bluebell_project/controller"
	"Gin_Project/src/18_bluebell_project/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
* Description:
* @Author Hollis
* @Create 2023/9/29 13:05
 */

func Setup() *gin.Engine {
	r := gin.Default()

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/signup", controller.SignUpHandler)

	// 登录业务路由
	r.POST("/login", controller.LoginHandler)

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
