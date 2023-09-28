package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/**
* Description:
	数据库连接的初始化操作
* @Author Hollis
* @Create 2023/9/27 21:52
*/
// 定义一个全局对象db
var db *sql.DB

// User 创建一个结构体和数据库表进行映射
type User struct {
	Id   int
	Name string
	Age  int
}

func init() {
	dsn := "root:abc123@tcp(127.0.0.1:3306)/gormdb"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		return
	}
	fmt.Println("数据库连接成功")
}
