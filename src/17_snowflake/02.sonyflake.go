package main

import (
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

/**
* Description:
	sony/sonyflake雪花算法生成id
* @Author Hollis
* @Create 2023/9/29 11:44
*/

// Settings startTime 选项和我们之前的 Epoch 差不多，如果不设置的话，默认是从2014-09-0100:00:00 +0000 UTC开始。
// MachineID 可以由用户自定义的函数，如果用户不定义的话，会默认将本机IP的低16位作为machine id。
// checkMachineID 是由用户提供的检查 MachineID 是否冲突的函数。
var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

// Init02 需传入当前的机器ID
func Init02(startTime string, machineId uint16) (err error) {
	sonyMachineID = machineId
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}
	settings := sonyflake.Settings{
		StartTime: st,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}
func GenID02() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("snoy flake not inited")
		return
	}
	id, err = sonyFlake.NextID() // 生成新的id
	return
}

func main() {
	if err := Init02("2020-07-01", 1); err != nil {
		fmt.Printf("Init failed, err:%v\n", err)
		return
	}
	for i := 0; i < 100; i++ {
		id, _ := GenID02()
		fmt.Println(id)
	}
}
