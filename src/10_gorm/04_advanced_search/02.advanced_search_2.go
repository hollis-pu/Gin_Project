package main

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
)

/*
*
  - Description:
    高级查询2：子查询、参数命名、find到map、查询引用Score
  - @Author Hollis
  - @Create 2023/9/15 14:51
*/
func main() {
	DB := DB.Session(&gorm.Session{Logger: mysqlLogger})
	//_ = DB.AutoMigrate(&Person{})
	// 1.子查询
	// 查询大于平均年龄的用户
	var users []Person
	DB.Model(Person{}).Where("age > (?)", DB.Model(Person{}).Select("avg(age)")).Find(&users)
	fmt.Println(users)
	// SELECT * FROM `people` WHERE age > (SELECT avg(age) FROM `people`)

	// 2.参数命名
	users = []Person{}
	// 方式1
	DB.Where("name = @name and age = @age", sql.Named("name", "枫枫"), sql.Named("age", 23)).Find(&users)
	// 方式2
	DB.Where("name = @name and age = @age", map[string]any{"name": "枫枫", "age": 23}).Find(&users)
	fmt.Println(users)
	// SELECT * FROM `people` WHERE name = '枫枫' and age = 23

	// 3.find到map
	var res []map[string]any
	DB.Table("people").Find(&res)
	// SELECT * FROM `people`
	fmt.Println(res)

	// 4.查询引用Scope：可以在model层写一些通用的查询方式，这样外界就可以直接调用方法即可
	users = []Person{}
	DB.Scopes(Age23).Find(&users)
	fmt.Println(users)

}

// Age23 定义一个函数，通过Scope进行调用
func Age23(db *gorm.DB) *gorm.DB {
	return db.Where("age > ?", 23)
}
