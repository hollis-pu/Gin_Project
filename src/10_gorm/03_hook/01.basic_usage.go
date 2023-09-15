package main

import (
	"fmt"
	"gorm.io/gorm"
)

/*
*
  - Description:
    hook的简单使用
  - @Author Hollis
  - @Create 2023/9/15 13:38
*/

func (p *Person) AfterFind(tx *gorm.DB) (err error) {
	fmt.Println("成功查询出一条记录！")
	return nil
}

func main() {
	DB := DB.Session(&gorm.Session{Logger: mysqlLogger})
	_ = DB.AutoMigrate(&Person{})

	var p1 Person
	DB.Take(&p1)
	fmt.Println(p1)
}
