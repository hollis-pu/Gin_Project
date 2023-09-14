package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
*
  - Description:
    普通路由的使用
  - @Author Hollis
  - @Create 2023/9/14 17:03
*/
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./404.html") // NoRoute()方法中用到了404.html，所以这里需要将该文件加载进来。
	// 访问/index的GET请求会走这一条处理逻辑
	// 路由
	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "GET",
		})
	})
	// 同样的方式创建其他 HTTP 方法的路由。

	// Any：可以匹配所有请求的方法。
	//r.Any("/test", func(c *gin.Context) {...})

	// 为没有配置处理函数的路由添加处理程序，默认情况下它返回404代码，下面的代码为没有匹配到路由的请求都返回views/404.html页面。
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	r.Run(":9006")
}
