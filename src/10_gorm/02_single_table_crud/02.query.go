package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

/**
* Description:
	Gorm进行单表查询的使用
* @Author Hollis
* @Create 2023/9/15 9:39
*/

func main() {
	DB := DB.Session(&gorm.Session{Logger: mysqlLogger})
	_ = DB.AutoMigrate(&Person{})

	// 3.查询单条记录：Take（注意：通过Take查询只能得到一条记录，即使条件查询的结果有多个，因为它会在条件的后面加上limit 1限制）
	// 3.1 取出一条记录
	var p2 Person
	DB.Take(&p2) // 默认排序，存入p2中
	// SELECT * FROM `people` LIMIT 1
	fmt.Println(p2) // {1 hollis 25 true 0xc0001f07f0}

	DB.First(&p2) // 按照主键升序的第一条记录
	// SELECT * FROM `people` WHERE `people`.`id` = 1 ORDER BY `people`.`id` LIMIT 1
	fmt.Println(p2) // {1 hollis 25 true 0xc0001ee8f0}

	DB.Last(&p2) // 按照主键降序的第一条记录
	// SELECT * FROM `people` WHERE `people`.`id` = 1 ORDER BY `people`.`id` DESC LIMIT 1
	fmt.Println(p2) // {1 hollis 25 true 0xc0001eea00}

	// 3.2 根据主键查询
	p2 = Person{} // 这里需要置空p2，否则会带上前面的查询条件
	DB.Take(&p2, 5)
	//  SELECT * FROM `people` WHERE `people`.`id` = 5 LIMIT 1
	fmt.Println(p2) // {5 机器人3号 7 true 0xc0001eeb20}

	// 3.3 根据其他条件查询
	p2 = Person{}                                                   // 这里需要置空p2，否则会带上前面的查询条件
	DB.Take(&p2, "name=? and email=?", "机器人7号", "robot7@gmail.com") // 使用？作为占位符，将查询的内容放入?  这样可以有效的防止sql注入，它的原理就是将参数全部转义。
	// SELECT * FROM `people` WHERE name='机器人7号' and email='robot7@gmail.com' LIMIT 1
	fmt.Println(p2) // {9 机器人7号 3 false 0xc0001eec30}

	// 4.查询多条记录：Find
	// 4.1 查询所有记录
	var p3 []Person
	DB.Find(&p3)
	// SELECT * FROM `people`
	for _, p := range p3 { // 可以遍历切片，得到每一条记录的值
		data, _ := json.Marshal(p) // 由于email是指针类型，所以看不到实际的内容。但是序列化之后，会转换为我们可以看得懂的方式。
		fmt.Println(string(data))
	}

	// 4.2 根据主键列表查询
	p3 = []Person{}
	//DB.Find(&p3, []int{3, 5, 7, 9})
	DB.Find(&p3, 3, 5, 7, 9) // 意思同上
	// SELECT * FROM `people` WHERE `people`.`id` IN (3,5,7,9)
	fmt.Println(p3)

	// 4.3 根据其他条件查询
	p3 = []Person{}
	DB.Find(&p3, "name in ?", []string{"hollis", "机器人3号", "机器人5号"})
	// SELECT * FROM `people` WHERE name in ('hollis','机器人3号','机器人5号')
	fmt.Println(p3)
}
