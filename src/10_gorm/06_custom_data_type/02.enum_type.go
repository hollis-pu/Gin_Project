package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

/*
*
  - Description:
    自定义数据类型——枚举
  - @Author Hollis
  - @Create 2023/9/16 10:46
*/

type Status int

const (
	Running Status = iota // 使用iota来递增常量
	Except
	OffLine
)

type Host struct {
	IP     string `json:"ip"`
	Status Status `json:"status"`
}

// MarshalJSON 实现MarshalJSON()方法，会在json.Marshal()时被自动调用
func (status Status) MarshalJSON() ([]byte, error) { // 作Status数字和字符串之间的映射
	var str string
	switch status {
	case Running:
		str = "Running"
	case Except:
		str = "Except"
	case OffLine:
		str = "OffLine"
	}
	return json.Marshal(str)
}

func main() {
	DB := DB.Session(&gorm.Session{Logger: mysqlLogger})
	_ = DB.AutoMigrate(&Host{}) // 自动迁移数据库结构

	// 添加数据
	DB.Create(&Host{
		IP:     "192.168.12.56",
		Status: Running,
	})
	DB.Create(&Host{
		IP:     "192.168.12.88",
		Status: OffLine,
	})

	// 查询数据
	var host Host
	DB.Take(&host, "ip = ?", "192.168.12.56")
	fmt.Println(host) // {192.168.12.56 1}
	// 这样打印出来的值是数据库中实际存放的值，如果我们想要给前端返回Status值而不是数字，就需要实现MarshalJSON()方法。
	// 在通过json.Marshal()对host进行序列化时，会自动调用MarshalJSON()函数，将数字转换为其对应的字符串。

	// 序列化为JSON格式的字符串
	data, _ := json.Marshal(host) // 自动调用MarshalJSON()函数
	fmt.Println(string(data))     // {"ip":"192.168.12.56","status":"Running"}
}
