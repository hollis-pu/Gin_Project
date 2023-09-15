package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
*
  - Description:
    中间件的简单使用
  - @Author Hollis
  - @Create 2023/9/14 19:48
*/

func indexHandler(c *gin.Context) {
	fmt.Println("indexHandler in...")
	name, ok := c.Get("name") // 获取上下文中保存的变量值（跨中间件存取值）
	if !ok {
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"username": name,
	})
}

// 定义一个中间件m1
func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	start := time.Now().UnixMicro()
	c.Set("name", "小王子") // 可以向上下文中写入变量，后续的中间件可以获取到
	// 调用该请求的剩余处理程序
	c.Next()

	// 不调用该请求的剩余处理程序
	//c.Abort()

	// 计算耗时
	cost := time.Now().UnixMicro() - start
	fmt.Printf("cost:%v\n", cost)
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if doCheck { // 如果doCheck==true，则进行检查，否则直接放行
			// 是否登录的判断
			// if 是登录用户
			// c.Next()
			// else
			// c.Abort()
		} else {
			c.Next()
		}
	}

}

func main() {
	r := gin.Default() // 默认使用了Logger和Recover中间件

	// 全局注册中间件函数m1  Use(middleware ...HandlerFunc)
	r.Use(m1, authMiddleware(false)) // 这样写了以后，每一个函数中就都会经过中间件m1,authMiddleware
	// GET(relativePath string, handlers ...HandlerFunc)
	//r.GET("/index", m1, indexHandler) // 先经过m1中间件，然后再经过indexHandler中间件

	r.GET("/index", indexHandler)

	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"content": "user",
		})
	})

	r.Run(":9001")
}
