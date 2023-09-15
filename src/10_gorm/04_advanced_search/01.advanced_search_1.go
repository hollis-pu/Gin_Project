package main

import (
	"fmt"
	"gorm.io/gorm"
)

/*
*
  - Description:
    高级查询内容，包括：where+find、Not、Or、Select、排序、分页、去重、分组和执行原生SQL语句。
  - @Author Hollis
  - @Create 2023/9/15 13:55
*/
func main() {
	DB := DB.Session(&gorm.Session{Logger: mysqlLogger})
	_ = DB.AutoMigrate(&Person{})

	var users []Person

	// 查询用户名是枫枫的
	DB.Where("name = ?", "枫枫").Find(&users)
	fmt.Println(users)

	// 查询用户名不是枫枫的
	DB.Where("name <> ?", "枫枫").Find(&users)
	fmt.Println(users)

	// 查询用户名包含 如燕，李元芳的
	DB.Where("name in ?", []string{"如燕", "李元芳"}).Find(&users)
	fmt.Println(users)

	// 查询姓李的
	DB.Where("name like ?", "李%").Find(&users)
	fmt.Println(users)

	// 查询年龄大于23，是qq邮箱的
	DB.Where("age > ? and email like ?", "23", "%@qq.com").Find(&users)
	fmt.Println(users)

	// 查询是qq邮箱的，或者是女的
	DB.Where("gender = ? or email like ?", false, "%@qq.com").Find(&users)
	fmt.Println(users)

	// 结构体查询：会过滤零值
	DB.Where(&Person{Name: "李元芳", Age: 0}).Find(&users)
	// SELECT * FROM `people` WHERE `people`.`name` = '李元芳'
	fmt.Println(users)

	// 使用map查询：不会过滤零值
	DB.Where(map[string]any{"name": "李元芳", "age": 0}).Find(&users)
	// SELECT * FROM `students` WHERE `age` = 0 AND `name` = '李元芳'
	fmt.Println(users)

	// Not条件
	// 排除年龄大于23的
	DB.Not("age > 23").Find(&users)
	fmt.Println(users)

	// Or条件：和where中的or等价
	DB.Or("gender = ?", false).Or(" email like ?", "%@qq.com").Find(&users)
	fmt.Println(users)

	// Select字段
	DB.Select("name", "age").Find(&users)
	fmt.Println(users) // 没有被选中的字段，会被赋零值（多出很多不需要的字段）

	// Scan：将选择的字段存入另一个结构体中（Scan是根据column列名进行扫描的）
	type User struct {
		Name string
		Age  int
	}
	var users01 []User
	DB.Model(&Person{}).Select("name", "age").Scan(&users01)
	//DB.Table("people").Select("name", "age").Scan(&users01)  // 和上面一句的功能相同
	fmt.Println(users01) // 这样得到的切片中，就只有需要的字段了

	// 排序
	var users02 []Person
	DB.Order("age desc").Find(&users02)
	fmt.Println(users02)
	// desc    降序
	// asc     升序

	// 分页
	var users03 []Person
	// 一页2条，第1页
	DB.Limit(2).Offset(0).Find(&users03)
	fmt.Println(users03)
	// 第2页
	DB.Limit(2).Offset(2).Find(&users03)
	fmt.Println(users03)
	// 第3页
	DB.Limit(2).Offset(4).Find(&users03)
	fmt.Println(users03)

	// 通用写法
	// 一页多少条
	limit := 2
	// 第几页
	page := 1
	offset := (page - 1) * limit
	DB.Limit(limit).Offset(offset).Find(&users)
	fmt.Println(users)

	// 去重
	var ageList []int
	DB.Table("people").Select("age").Distinct("age").Scan(&ageList)
	// SELECT DISTINCT age FROM `people`
	fmt.Println(ageList)

	// 分组查询
	type AggeGroup struct {
		Gender int
		Count  int    `gorm:"column:count(id)"`
		Name   string `gorm:"column:group_concat(name)"`
	}

	var agge []AggeGroup
	// 查询男生的个数和女生的个数，再将名字进行汇总
	DB.Table("people").Select("count(id)", "gender", "group_concat(name)").Group("gender").Scan(&agge)
	// SELECT count(id),gender,group_concat(name) FROM `people` GROUP BY `gender`
	fmt.Println(agge) // [{0 3 李琦,晓梅,如燕} {1 6 李元芳,张武,枫枫,刘大,李武,魔灵}]

	// 执行原生sql
	var peronList []Person
	DB.Raw(`SELECT * FROM people where age=23`).Find(&peronList)
	fmt.Println(peronList)

}
