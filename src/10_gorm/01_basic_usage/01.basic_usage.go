package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/*
*
  - Description:
    Gorm的简单使用：建表、命名策略、关系映射、设置日志、gorm标签等的使用。
  - @Author Hollis
  - @Create 2023/9/15 09:01
*/

type Student struct {
	ID    uint    // 默认使用ID作为主键
	Name  string  `gorm:"size:20;not null:true"`
	Age   int     `gorm:"size:4"`
	Email *string `gorm:"size:32"` // 使用指针是为了存空值
}

//func (Student) TableName() string { // 指定表名
//	return "student"
//}

var DB *gorm.DB
var mysqlLogger logger.Interface

func init() {
	username := "root"   //账号
	password := "abc123" //密码
	host := "127.0.0.1"  //数据库地址，可以是Ip或者域名
	port := 3306         //数据库端口
	Dbname := "gormdb"   //数据库名
	timeout := "10s"     //连接超时，10秒

	mysqlLogger = logger.Default.LogMode(logger.Info) // 要显示的日志等级（低到高：Silent、Error、Warn、Info）

	// root:root@tcp(127.0.0.1:3306)/gormdb?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		//NamingStrategy: schema.NamingStrategy{ // 配置表名转换的规则
		//	TablePrefix:   "f_", // 表名前缀
		//	SingularTable: true, // 是否单数表名
		//	NoLowerCase:   true, // 是否小写转换
		//},
		//Logger: mysqlLogger, // 全局显示日志
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 连接成功
	DB = db
}

func main() {

	DB := DB.Session(&gorm.Session{Logger: mysqlLogger}) // 为DB设置Logger

	err := DB.AutoMigrate(&Student{}) // 创建数据库表students和结构体Student的映射关系
	if err != nil {
		return
	}
	/* 创建的表结构如下
	create table students
	(
	    id    bigint unsigned auto_increment
	        primary key,
	    name  varchar(20) not null,
	    email varchar(32) null,
	    age   tinyint     null
	);
	*/
}
