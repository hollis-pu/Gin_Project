package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
*
  - Description:
    获取form表单提交的数据：当前端请求的数据通过form表单提交时，例如向/user/search发送一个POST请求。
  - @Author Hollis
  - @Create 2023/9/14 10:58
*/
func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./login.html", "./index.html")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// 接收前端通过form的submit发送到后端的post请求
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username") // 获取post请求中form表单中的元素内容
		password := c.PostForm("password")
		c.HTML(http.StatusOK, "index.html", gin.H{ // 跳转到index页面
			"Name":     username,
			"Password": password,
		})
	})

	r.Run(":9005")
}
