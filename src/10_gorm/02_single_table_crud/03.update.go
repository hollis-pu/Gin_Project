package main

import (
	"fmt"
	"gorm.io/gorm"
)

/**
* Description:
	Gorm进行更新数据的操作
* @Author Hollis
* @Create 2023/9/15 9:39
*/

func main() {
	DB := DB.Session(&gorm.Session{Logger: mysqlLogger})
	_ = DB.AutoMigrate(&Person{})

	// 5.更新：Save
	// 5.1 更新全部字段
	var p4 Person
	err := DB.Take(&p4, "name=?", "机器人5号").Error // 更新之前先查询记录，找到记录的id（主键）
	if err != nil {
		fmt.Println("查询记录失败！")
		panic(err)
	}
	p4.Age = 10
	p4.Gender = false
	DB.Save(&p4) // 全字段更新
	// UPDATE `people` SET `name`='机器人5号',`age`=10,`gender`=false,`email`='robot5@gmail.com' WHERE `id` = 7

	// 5.2 更新指定字段：Select选定字段 Omit忽略字段
	p4 = Person{}
	DB.Take(&p4, "name=?", "机器人5号")
	p4.Age = 20
	p4.Gender = true
	DB.Select("age", "gender").Save(&p4) // 使用Select选择要更新的字段
	fmt.Println(p4)                      // {7 机器人5号 20 true 0xc0001eeb00}

	// 6.批量更新
	// 6.1 更新单个字段
	var p5 []Person
	DB.Find(&p5, "age=?", 12).Update("age", 12) // 先保存查询的结果集的id到切片p5中，然后执行批量更新的操作
	// SELECT * FROM `people` WHERE age=12
	// UPDATE `people` SET `age`=12 WHERE age=12 AND `id` IN (6,8,11,12,13)

	// 6.2 更新多个字段
	p5 = []Person{}
	// 方式1：结构体（当使用struct更新时，GORM只会更新非零字段。如果想要更新非零字段，请用于map更新属性或使用Select指定要更新的字段）
	DB.Find(&p5, "age=? and gender=?", 12, false).Select("age", "gender").Updates(Person{ // 加上Select指定要跟心的字段
		Age:    12,
		Gender: false, // 默认不会更新字段为0(false)的值，需要在前面加上Select才行
	})
	// SELECT * FROM `people` WHERE age=12 and gender=false
	// UPDATE `people` SET `age`=12,`gender`=false WHERE (age=12 and gender=false) AND `id` IN (8,11)

	// 方式2：map（效果同方式1）
	DB.Find(&p5, "age=? and gender=?", 12, false).Updates(map[string]any{
		"age":    12,
		"gender": false,
	})
}
