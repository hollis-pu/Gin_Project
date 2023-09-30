package router

import (
	"Gin_Project/src/18_bluebell_project/controller"
	"Gin_Project/src/18_bluebell_project/logger"
	"Gin_Project/src/18_bluebell_project/middlewares"
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

	r.GET("/ping", func(c *gin.Context) { // 测试路由能否正常访问
		c.String(http.StatusOK, "pong")
	})

	// 注册业务路由
	r.POST("/signup", controller.SignUpHandler)

	// 登录业务路由
	r.POST("/login", controller.LoginHandler)

	// 现在对用户的访问进行限制，只有已经登录的用户才能访问"/ping"页面
	r.GET("/isLogin", middlewares.JWTAuthMiddleware(), func(c *gin.Context) { // 加上JWTAuthMiddleware认证的中间件
		// 只要中间件能够正常执行完成，则说明用户已经登录
		c.String(http.StatusOK, "ok!")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
