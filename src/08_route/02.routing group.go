package main

import "github.com/gin-gonic/gin"

/*
*
  - Description:
    路由组的基本使用
  - @Author Hollis
  - @Create 2023/9/14 18:07
*/
func main() {
	r := gin.Default()

	hollis := r.Group("hollis") // 创建一个路由组，作为公共前缀
	// 可以给属于这个组的处理函数外面加上一个大括号，看上去更清晰。
	{
		hollis.GET("/index", func(c *gin.Context) { // 访问时需访问：http://localhost:9001/hollis/index
			c.JSON(200, gin.H{
				"page": "index",
			})
		})
		hollis.GET("/picture", func(c *gin.Context) { // 访问时需访问：http://localhost:9001/hollis/index
			c.JSON(200, gin.H{
				"page": "picture",
			})
		})
		hollis.GET("/blog", func(c *gin.Context) { // 访问时需访问：http://localhost:9001/hollis/index
			c.JSON(200, gin.H{
				"page": "blog",
			})
		})
	}
	r.Run(":9001")
}
