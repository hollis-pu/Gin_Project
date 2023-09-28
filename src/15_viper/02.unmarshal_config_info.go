package main

import (
	"fmt"
	"github.com/spf13/viper"
)

/*
*
  - Description:
    Unmarshal：将配置文件中所有或特定的值解析到结构体、map等
  - @Author Hollis
  - @Create 2023/9/28 18:00
*/

type Config struct {
	Port        int    `mapstructure:"port"`
	Version     string `mapstructure:"version"`
	MysqlConfig `mapstructure:"mysql"`
}

type MysqlConfig struct {
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	Dbname string `mapstructure:"dbname"`
}

func main() {
	viper.SetConfigFile("./src/15_viper/config.yaml") // 指定配置文件路径

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var c Config
	err = viper.Unmarshal(&c)
	if err != nil {
		fmt.Printf("Unmarshal config failed, err:%v\n", err)
		return
	}
	fmt.Printf("%#v\n", c)
}
