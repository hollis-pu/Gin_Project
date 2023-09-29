package main

import (
	"fmt"
	sf "github.com/bwmarrin/snowflake"
	"time"
)

/**
* Description:
	bwmarrin/snowflake雪花算法生成id
* @Author Hollis
* @Create 2023/9/29 10:55
*/

var node *sf.Node

func Init01(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}

func GenID01() int64 {
	return node.Generate().Int64()
}
func main() {
	if err := Init01("2020-07-01", 1); err != nil {
		fmt.Printf("init failed, err:%v\n", err)
		return
	}
	for i := 0; i < 100; i++ { // 生成100个id
		id := GenID01()
		fmt.Println(id)
	}
}
