package main

import "gorm.io/gorm"

/*
*
  - Description:
    一对多的删除：级联删除和清除外键关系
  - @Author Hollis
  - @Create 2023/9/15 21:07
*/
func main() {
	DB := DB.Session(&gorm.Session{Logger: mysqlLogger})
	_ = DB.AutoMigrate(&User{})
	_ = DB.AutoMigrate(&Article{})

	//这里先造一些测试用的数据
	DB.Create(&User{
		Name: "test1",
		Articles: []Article{
			{
				Title: "test1 article01",
			},
			{
				Title: "test1 article02",
			},
		},
	})
	DB.Create(&User{
		Name: "test2",
		Articles: []Article{
			{
				Title: "test2 article01",
			},
			{
				Title: "test2 article02",
			},
		},
	})

	// 当删除文章时，对users表不会有任何的影响。
	// 但是当删除用户时，就需要考虑是否删除articles表中的数据了。
	//
	// 清除外键关系
	// 方案1：删除用户时，不删除articles表中的记录，而是将其中的uid字段设置为null。
	// 这个方案需要先清除外键关系，然后再删除users表中的用户信息
	var user User
	DB.Preload("Articles").Take(&user, "name =?", "test1")             // 找到指定用户
	_ = DB.Model(&user).Association("Articles").Delete(&user.Articles) // 删除外键关系
	DB.Delete(&user)
	// 这样，既删除了用户，也将articles表中的uid字段设置为了null。

	// 级联删除
	// 方案2：删除用户时，同时删除articles表中的记录。
	var user02 User
	DB.Take(&user02, "name = ?", "test2") // 找到指定用户
	DB.Select("Articles").Delete(&user02) // 选中需要删除的Articles字段，然后直接删除用户，这样articles表相应的记录也会被删除。
	// 这样，在删除用户的同时，也删除了外键关联的文章记录。
}
