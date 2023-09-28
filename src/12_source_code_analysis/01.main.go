package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
* Description:
* @Author Hollis
* @Create 2023/9/16 21:07
 */
func handler01(c *gin.Context) {
	fmt.Println("handler01")
	c.Set("cookie", "123234")
}

func handler02(c *gin.Context) {
	fmt.Println("handler01")
	cookie, _ := c.Get("cookie")
	fmt.Println(cookie)
}

func main() {
	r := gin.Default()

	groupRoute := r.Group("/test")
	groupRoute.Use(handler01, handler02)
	groupRoute.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	_ = r.Run()
}
