package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
*
  - Description:
    使用Gin框架，渲染从网上下载的网页模板
    网页模板下载地址：https://sc.chinaz.com/moban/230911032070.htm
  - @Author Hollis
  - @Create 2023/9/13 23:26
*/
func main() {
	r := gin.Default()
	r.Static("static", "./statics") // 加载静态资源。这里是指定静态文件的路径映射。模板文件中所有以static开头的静态文件都到当前目录下的statics目录去找。
	r.LoadHTMLGlob("*.html")        // 加载所有的html文件
	//r.LoadHTMLFiles("index.html")

	r.GET("/404.html", func(c *gin.Context) { // 当访问http://localhost:9001/index.html时，返回指定的模板文件。
		c.HTML(http.StatusOK, "404.html", nil)
	})
	r.GET("/about1.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about1.html", nil)
	})
	r.GET("/about2.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about2.html", nil)
	})
	r.GET("/blog.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "blog.html", nil)
	})
	r.GET("/blog-single.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "blog-single.html", nil)
	})
	r.GET("/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/portfolio.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "portfolio.html", nil)
	})
	r.GET("/services.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "services.html", nil)
	})

	r.Run(":9001") // 指定监听的端口
}
