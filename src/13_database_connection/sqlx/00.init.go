package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 注册驱动
	"github.com/jmoiron/sqlx"
)

/**
* Description:
	数据库连接的初始化操作
* @Author Hollis
* @Create 2023/9/27 21:52
*/
// 定义一个全局对象db
var db *sqlx.DB

// User 创建一个结构体和数据库表进行映射
type User struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func init() {
	dsn := "root:abc123@tcp(127.0.0.1:3306)/gormdb"
	var err error
	db, err = sqlx.Connect("mysql", dsn) // Connect=原生sql包中的Open+Ping
	// 也可以使用MustConnect连接不成功就panic
	//db= sqlx.MustConnect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
}
