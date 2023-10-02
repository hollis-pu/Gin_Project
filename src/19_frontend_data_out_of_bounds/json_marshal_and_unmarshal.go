package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type MyData struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

/*
*
  - Description:
  - @Author Hollis
  - @Create 2023/10/2 20:46
*/
func main() {
	data1 := MyData{
		ID:   math.MaxInt64,
		Name: "Tom",
	}

	// 序列化
	b, err := json.Marshal(data1)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(string(b))

	// 反序列化
	s := `{"id":9223372036854775807,"name":"Jerry"}`
	var data2 MyData
	if err := json.Unmarshal([]byte(s), &data2); err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(data2)

}
