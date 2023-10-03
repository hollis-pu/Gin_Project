package main

import (
	"Gin_Project/src/22_limit_traffic/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/**
* Description:
* @Author Hollis
* @Create 2023/10/3 18:07
 */
func main() {
	r := gin.Default()

	v1 := r.Group("/v1", middlewares.RateLimitMiddleware1()) // 每1秒可以访问1次
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "ok",
		})
	})

	v2 := r.Group("/v2", middlewares.RateLimitMiddleware2(2*time.Second, 1)) // 每两秒可以访问1次
	v2.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "ok",
		})
	})
	r.Run()
}
