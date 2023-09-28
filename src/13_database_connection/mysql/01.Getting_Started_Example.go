package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 匿名导入，自动执行mysql包中的init()
	"time"
)

/**
* Description:
	Go语言连接mysql数据库的入门案例
* @Author Hollis
* @Create 2023/9/27 15:51
*/

func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:abc123@tcp(127.0.0.1:3306)/atguigu"

	// 注意Open()不会校验账号密码是否正确
	db, err := sql.Open("mysql", dsn) // 这里不能保证连接成功
	if err != nil {
		panic(err)
	}

	// 使用Ping()来确保数据库连接成功
	err = db.Ping()
	if err != nil {
		return err
	}
	// 设置参数
	db.SetMaxOpenConns(100)                // 最大连接数
	db.SetMaxIdleConns(10)                 // 最大空闲连接数
	db.SetConnMaxLifetime(time.Minute * 3) //连接的最大生存时间
	return nil
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	fmt.Println("init db successful!")
}
