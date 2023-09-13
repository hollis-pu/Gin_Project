package main

import "github.com/gin-gonic/gin"

/*
*
  - Description:
    RESTful API的简单使用
  - GET用来获取资源；
  - POST用来新建资源；
  - PUT用来更新资源；
  - DELETE用来删除资源。
    只要API程序遵循了REST风格，那就可以称其为RESTful API。
    可以使用Postman来对put和delete方法进行测试。
  - @Author Hollis
  - @Create 2023/9/13 16:58
*/
func main() {
	r := gin.Default()

	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "GET",
		})

	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "POST",
		})

	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "PUT",
		})

	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "DELETE",
		})
	})

	r.Run(":9090")
}
