package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

/**
* Description:
* @Author Hollis
* @Create 2023/9/29 13:16
 */

// Conf 全局变量，用来保存程序的所有配置信息
var Conf = new(AppConfig) // 这里创建的是一个指针，保证在修改的时候是操作的同一个变量

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	StartTime    string `mapstructure:"start_time"` // 雪花算法生成ID的起始时间
	MachineID    int64  `mapstructure:"machine_id"` // 雪花算法生成ID的工作机器id
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host        string `mapstructure:"host"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Dbname      string `mapstructure:"dbname"`
	Port        int    `mapstructure:"port"`
	MaxOpenConn int    `mapstructure:"max_open_conns"`
	MaxIdleConn int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Read config file failed, err:%v\n", err)
		return
	}
	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) { // 回调
		fmt.Println("Config file changed:", e.Name)
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return
}
