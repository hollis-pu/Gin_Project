package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
*
  - Description:
    获取路径参数的简单使用
  - @Author Hollis
  - @Create 2023/9/14 13:11
*/
func main() {
	r := gin.Default()

	// 注意URL的匹配不要冲突
	r.GET("blog/:year/:month", func(c *gin.Context) { // 假设这里将博客的文章按照年份和月份设置path
		year := c.Param("year")
		month := c.Param("month")

		c.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})

	r.Run(":9006")
}
