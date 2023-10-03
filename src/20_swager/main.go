package main

import (
	"Gin_Project/src/20_swager/controller"
	_ "Gin_Project/src/20_swager/docs" // 千万不要忘了导入把你上一步生成的docs
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

/**
* Description:
	swagger生成API文档的使用
* @Author Hollis
* @Create 2023/10/3 10:26
*/

//	@title			Swagger接口文档
//	@version		1.0
//	@description	这里写描述信息
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	这里写联系人信息
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", controller.PingHandler)
		v1.POST("/login", controller.LoginHandler)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.Run()
}
