package main

import "gorm.io/gorm"

/*
*
  - Description:
    表的关联：一对多的关系（添加）
  - @Author Hollis
  - @Create 2023/9/15 16:31
*/
func main() {
	DB := DB.Session(&gorm.Session{Logger: mysqlLogger})
	_ = DB.AutoMigrate(&User{})
	_ = DB.AutoMigrate(&Article{})

	// 1.一对多的添加
	// 1.1 创建用户，并且创建文章
	a1 := Article{Title: "python"}
	a2 := Article{Title: "golang"}
	user := User{Name: "hollis", Articles: []Article{a1, a2}} // User结构体中拥有了Article结构体
	DB.Create(&user)

	// 1.2 创建文章，关联已有用户
	// 方式1：指定UID
	a3 := Article{Title: "redis", UID: 1} // 创建文章对象时，指定UID
	DB.Create(&a3)

	// 方式2：查询一个user，传入Article结构体变量中
	var user01 User
	DB.Take(&user01, 1) // 获取一个User对象
	DB.Create(&Article{
		Title: "how to learn golang?",
		User:  user01, // 这里也可以传入一个User对象，效果同上
	})

	// 1.3 给已有的用户绑定文章
	// 方式1：指定属性值
	var user02 User
	DB.Take(&user02, "name = ?", "hollis")
	var article01 Article
	DB.Take(&article01, 3)

	article01.User = user02 // 设置指定article的user
	DB.Save(&article01)

	// 方式2：Association+Append
	_ = DB.Model(&user02).Association("Articles").Append(&article01) // 效果同上

}
