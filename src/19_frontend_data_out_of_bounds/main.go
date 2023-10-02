package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Data struct {
	ID int64 `json:"id,string"` // 在json标签中加入string即可
}

/*
*
  - Description:
    在我们使用雪花算法生成id时，通常会生成类型为int64的id，但是前端在处理数据时，大概率会使用JavaScript来进行处理。
    Go语言中int64的表示范围：-2^62---2^62(包含边界)
    但是JS中整数类型的范围为： -2^53---2^53(包含边界)。
    如何解决前端数据越界的问题？将可能越界的字段转换为string进行序列化和反序列化。
  - @Author Hollis
  - @Create 2023/10/2 21:20
*/
func main() {
	r := gin.Default()

	r.LoadHTMLFiles("index.html")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/data", func(c *gin.Context) {
		c.JSON(http.StatusOK, Data{
			9223372036854775807,
			// 9223372036854776000（前端js中保存的值，造成了数据越界）
			// "9223372036854775807"（在json tag中加入string之后，js中保存的值）
		})
	})

	r.Run(":8090")
}
