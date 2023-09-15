package main

import (
	"fmt"
	"gorm.io/gorm"
)

/*
*
  - Description:
    一对多的查询
  - @Author Hollis
  - @Create 2023/9/15 20:37
*/
func main() {
	DB := DB.Session(&gorm.Session{Logger: mysqlLogger})
	_ = DB.AutoMigrate(&User{})
	_ = DB.AutoMigrate(&Article{})

	var article01 Article
	DB.Find(&article01, 1)
	// SELECT * FROM `articles` WHERE `articles`.`id` = 1
	fmt.Println(article01) // {1 python 1 {0  []}}  这样只能拿到UID，不能拿到用户的具体信息

	// 要拿到关联表的信息，就需要用到预加载
	// 预加载的名字是外键关联的属性名
	article01 = Article{}
	DB.Preload("User").Find(&article01, 1) // 这里的User为Article结构体中的User字段
	// SELECT * FROM `users` WHERE `users`.`id` = 1
	// SELECT * FROM `articles` WHERE `articles`.`id` = 1
	fmt.Println(article01) // {1 python 1 {1 hollis []}}

	// 也可以嵌套预加载，如查询出用户及其发布的文章（文章中也要包含User信息）
	var user01 User
	DB.Preload("Articles.User").Find(&user01, 1)
	fmt.Println(user01)

	// 带条件的预加载
	var user02 User
	DB.Preload("Articles", "id > ?", 2).Find(&user02, 1) // 只预加载Articles id大于2的文章
	fmt.Println(user02)

	// 自定义预加载
	var user03 User
	DB.Preload("Articles", func(db *gorm.DB) *gorm.DB { // 这里写一个自定义函数，函数签名为：func(db *gorm.DB) *gorm.DB
		return db.Where("id in ?", []int{2, 3})
	}).Find(&user03, 1) // 只预加载Articles id大于2的文章
	fmt.Println(user03)
}
