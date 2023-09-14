package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

/*
*
  - Description:
    http重定向和路由重定向的简单使用
  - @Author Hollis
  - @Create 2023/9/14 15:44
*/
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html", "./success.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 接收 action=/upload 的post请求
	r.POST("/upload", func(c *gin.Context) {
		// 从请求中读取文件
		file, err := c.FormFile("f1") // 这里的f1就是html文件中，file输入框的name。（和从请求中获取携带的参数一样）
		//file, err := c.MultipartForm() // 多个文件上传
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		} else {
			filePath := path.Join("./", file.Filename)           // 指定文件保存的路径
			_ = c.SaveUploadedFile(file, filePath)               // 将读取到的文件保存
			c.HTML(http.StatusOK, "success.html", file.Filename) // 跳转到success.html页面
		}
	})
	r.Run(":9009")
}
