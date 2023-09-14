package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
*
  - Description:
    获取querystring参数的简单使用：
  - @Author Hollis
  - @Create 2023/9/14 10:30
*/
func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {

		// Get请求 URL?后面的是querystring参数
		// key=value格式，多个key=value用 & 连接
		// 如：web?name=tom&age=19

		// 获取浏览器那边发送请求时携带的querystring
		// name := c.Query("name") // 获取参数name传递过来的值。如：前端请求http://localhost:9004/web?name=Tony，就可以获取到Tony
		// name := c.DefaultQuery("name", "default_value") // 如果能查到该参数，则使用该参数值。否则使用默认值。
		name, ok := c.GetQuery("name") // 和Query("name")不同的是，这里会多返回一个参数是否获取成功的bool值。
		if !ok {                       // 如果取不到参数值，就使用默认值
			name = "default_value"
		}
		// 也可以获取多个参数
		age := c.DefaultQuery("age", "0") // 设置默认值为0

		// 请求连接：http://localhost:9004/web?name=tom&age=19
		// 返回结果：{"age":"19","name":"tom"}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run(":9004")
}
