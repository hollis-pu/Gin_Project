package main

import (
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
)

/**
* Description:
	Gorm进行插入数据的操作
* @Author Hollis
* @Create 2023/9/15 9:39
*/

func main() {
	DB := DB.Session(&gorm.Session{Logger: mysqlLogger})
	_ = DB.AutoMigrate(&Person{})

	//1.插入单条记录
	email := "hollis_pu@163.com"
	p1 := Person{
		Name:   "hollis",
		Age:    25,
		Gender: true,
		Email:  &email,
	}
	DB.Create(&p1) // 创建一条记录（每执行一次，就会向数据库中写入一条数据）
	//指针类型是为了更好的存null类型，但是传值的时候，也记得传指针。
	//Create接收的是一个指针，而不是值。

	// 2.批量插入记录
	genderOpt := []bool{true, false}
	var personList []Person // 定义一个切片，来存放需要批量插入的数据
	for i := 0; i < 20; i++ {
		emails := fmt.Sprintf("robot%s@gmail.com", strconv.Itoa(i+1))
		personList = append(personList, Person{
			Name:   fmt.Sprintf("机器人%d号", i+1),
			Age:    rand.Intn(16) + 1,
			Gender: genderOpt[rand.Intn(2)], // 随机设置性别
			Email:  &emails,
		})
	}
	DB.Create(&personList) // 通过切片来批量插入
}
