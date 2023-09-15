package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/**
* Description:
	init函数，包内共享。用于创建结构体和连接数据库的操作。
* @Author Hollis
* @Create 2023/9/15 16:33
*/

var DB *gorm.DB
var mysqlLogger logger.Interface

// 一对多的关系。一个用户可以发布多篇文章，一篇文章属于一个用户

type User struct {
	ID       uint      `gorm:"size:4"`
	Name     string    `gorm:"size:8"`
	Articles []Article `gorm:"foreignKey:UID"` // 用户拥有的文章列表
}

type Article struct {
	ID    uint   `gorm:"size:4"`
	Title string `gorm:"size:120"`
	UID   uint   `gorm:"size:4"`         // 属于   这里的类型要和引用的外键类型一致，包括大小
	User  User   `gorm:"foreignKey:UID"` // 属于   这里重写了外键关联
}

func init() {
	username := "root"   //账号
	password := "abc123" //密码
	host := "127.0.0.1"  //数据库地址，可以是Ip或者域名
	port := 3306         //数据库端口
	Dbname := "gormdb"   //数据库名
	timeout := "10s"     //连接超时，10秒

	mysqlLogger = logger.Default.LogMode(logger.Info)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 连接成功
	DB = db
}
