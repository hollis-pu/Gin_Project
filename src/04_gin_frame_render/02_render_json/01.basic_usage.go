package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
*
  - Description:
    Gin框架渲染json数据的简单使用
  - @Author Hollis
  - @Create 2023/9/14 9:45
*/
func main() {
	r := gin.Default()

	// 方法1：使用map
	r.GET("/json01", func(c *gin.Context) {
		data := gin.H{ // gin.H is a shortcut for map[string]any
			"name":   "tom",
			"age":    26,
			"gender": "male",
		}
		c.JSON(http.StatusOK, data)
	})

	// 方法2：使用struct，使用tag来定制json字段的名称。
	r.GET("/json02", func(c *gin.Context) {
		var person struct {
			Name   string `json:"name"` // 这里一定要是可导出的，否则json获取不到数据。如果想要指定json字段的名称，可以使用Tag。
			Age    int    `json:"age"`
			Gender string `json:"gender"`
		}
		person.Name = "Jane"
		person.Age = 24
		person.Gender = "female"
		c.JSON(http.StatusOK, person) // 传入结构体类型的数据
	})

	r.Run(":9003")
}
