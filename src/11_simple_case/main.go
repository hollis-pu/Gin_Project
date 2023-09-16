package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

/**
* Description:
* @Author Hollis
* @Create 2023/9/16 16:23
 */

var DB *gorm.DB

func initMysql() {
	username := "root"
	password := "abc123"
	host := "127.0.0.1"
	port := 3306
	dbname := "gormdb"
	timeout := "10s"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		username, password, host, port, dbname, timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	DB = db
}

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func main() {
	initMysql()

	err := DB.AutoMigrate(&Todo{})
	if err != nil {
		fmt.Println("数据库迁移失败，err=", err)
		return
	}

	route := gin.Default()
	route.Static("/static", "static")
	route.LoadHTMLGlob("templates/*")

	// 进入主页
	route.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := route.Group("/v1") // 路由分组，给URL添加前缀
	{
		// 获取所有todo list
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []Todo
			DB.Find(&todoList)              // 查询出所有的结果集
			c.JSON(http.StatusOK, todoList) // 返回给前端
		})

		// 新增
		v1Group.POST("/todo", func(c *gin.Context) {
			var todo Todo
			_ = c.BindJSON(&todo)       // 绑定JSON的参数（将拿到的json数据传递给结构体变量todo）
			DB.Create(&todo)            // 写入数据库
			c.JSON(http.StatusOK, todo) // 返回给前端
		})
		// 修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			var todo Todo
			id := c.Param("id")                // 获取路径参数值
			DB.Where("id =?", id).First(&todo) // 查询指定id的第一条记录
			_ = c.BindJSON(&todo)              // 绑定Json的参数
			DB.Save(&todo)                     // 修改数据库
			c.JSON(http.StatusOK, todo)        // 返回给前端
		})

		// 删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id := c.Param("id")
			var todo Todo
			DB.Take(&todo, id)
			DB.Delete(&todo)
			c.JSON(http.StatusOK, gin.H{id: "deleted"})
		})
	}

	_ = route.Run(":8090")

}
