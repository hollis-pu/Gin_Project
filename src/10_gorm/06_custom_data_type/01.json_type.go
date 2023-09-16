package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

/**
* Description:
	自定义数据类型——Json：必须要实现Value()和Scan()方法
* @Author Hollis
* @Create 2023/9/16 10:21
*/

type Item struct {
	ID   uint
	Data MyJSON // 使用自定义JSON数据类型
}

// MyJSON 自定义JSON类型
type MyJSON map[string]interface{}

// Value 实现Valuer接口，将MyJSON类型转换为数据库字段值（存入数据库）
func (mj MyJSON) Value() (driver.Value, error) {
	return json.Marshal(mj)
}

// Scan 实现Scanner接口，从数据库字段值扫描到MyJSON类型（从数据库中取出）
func (mj *MyJSON) Scan(value interface{}) error {
	if value == nil {
		*mj = nil
		return nil
	}

	// 将数据库中的JSON字符串解析为MyJSON类型
	var jsonData []byte
	switch val := value.(type) {
	case []byte:
		jsonData = val
	case string:
		jsonData = []byte(val)
	default:
		return errors.New("failed to scan MyJSON from database")
	}

	return json.Unmarshal(jsonData, mj)
}

func main() {
	DB := DB.Session(&gorm.Session{Logger: mysqlLogger})
	_ = DB.AutoMigrate(&Item{}) // 自动迁移数据库结构

	// 创建一条记录
	DB.Create(&Item{
		Data: MyJSON{
			"product_name": "golang from entry to proficiency",
			"price":        89,
			"created_time": time.Now().Format("2006-01-02 15:04:05"),
		},
	})

	// 查询
	var item Item
	DB.Take(&item, 3)
	fmt.Println(item)
}
