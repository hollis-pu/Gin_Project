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
* @Create 2023/9/15 13:39
*/

var DB *gorm.DB
var mysqlLogger logger.Interface

type Person struct {
	ID     uint    `gorm:"size:10" json:"id"`
	Name   string  `gorm:"size:10;not null:true" json:"name"`
	Age    int     `gorm:"size:3" json:"age,omitempty"`
	Gender bool    `json:"gender"`
	Email  *string `gorm:"size:32" json:"email"`
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
