package main

import (
	"gorm.io/gorm"
)

/**
* Description:
	Gorm进行删除数据的操作
* @Author Hollis
* @Create 2023/9/15 9:39
*/

func main() {
	DB := DB.Session(&gorm.Session{Logger: mysqlLogger})
	_ = DB.AutoMigrate(&Person{})

	// 7.删除
	var p6 Person
	// 7.1 删除一条记录
	p6.ID = 12
	DB.Delete(&p6)
	// DELETE FROM `people` WHERE `people`.`id` = 12

	// 7.1 删除多条记录
	DB.Delete(&p6, 14, 15)
	// DELETE FROM `people` WHERE `people`.`id` IN (14,15)

	// 删除查询到的切片列表
	var p7 []Person
	DB.Find(&p7, "name in ?", []string{"机器人19号", "机器人20号"}).Delete(&p7)
	// SELECT * FROM `people` WHERE name in ('机器人19号','机器人20号')
	// DELETE FROM `people` WHERE name in ('机器人19号','机器人20号') AND `people`.`id` IN (24,25)
}
