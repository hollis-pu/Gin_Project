package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
*
  - Description:
    http重定向
  - @Author Hollis
  - @Create 2023/9/14 16:12
*/
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./redirect.html") // 这里不要忘记加载html文件
	r.GET("/index", func(c *gin.Context) {
		// 外部重定向：重定向到 baidu
		c.Redirect(http.StatusFound, "https://www.baidu.com/")
	})

	r.GET("/test", func(c *gin.Context) {
		// 内部重定向：跳转到 /redirect 对应的路由处理函数
		c.Request.URL.Path = "/redirect" // 把请求的URI修改
		r.HandleContext(c)               // 继续后续的处理
	})

	r.GET("/redirect", func(c *gin.Context) {
		c.HTML(http.StatusOK, "redirect.html", nil) // 2023.09.14 重定向之后的返回结果是html时，会报错，但是不知道原因。
		// 通过求助Chatgpt，已经知道的原因。是因为我没有将html文件导入到程序中，需要加上：r.LoadHTMLFiles("./redirect.html")。
		//c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	r.Run(":9090")
}
