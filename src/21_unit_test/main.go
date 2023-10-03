package main

import (
	"Gin_Project/src/21_unit_test/controller"
	"github.com/gin-gonic/gin"
)

/**
* Description:
* @Author Hollis
* @Create 2023/10/3 16:50
 */
func main() {
	r := gin.Default()

	r.GET("/ping", controller.PingHandler)
	r.POST("/login", controller.LoginHandler)

	r.Run()
}
