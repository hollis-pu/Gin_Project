package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

/*
*
  - Description:
    viper：读取配置文件，并实时监控配置文件的变化。
  - @Author Hollis
  - @Create 2023/9/28 15:41
*/
func main() {
	viper.SetDefault("version", "v0.0.0") // 设置默认值

	viper.SetConfigFile("./src/15_viper/config.yaml") // 指定配置文件路径
	// 也可以采用下面的方式：
	// viper.SetConfigName("config")         // 配置文件名称(无扩展名)
	// viper.SetConfigType("yaml")           // 如果配置文件的名称中没有扩展名，则需要配置此项
	// viper.AddConfigPath("./")             // 查找配置文件所在的路径
	// viper.AddConfigPath("./src/15_viper") // 多次调用以添加多个搜索路径

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 实时监控配置文件的变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})

	r := gin.Default()

	// 覆盖设置：优先级高于配置文件，会覆盖掉配置文件中的值
	//viper.Set("version", "v0.1.0")

	// 注册和使用别名
	viper.RegisterAlias("loud", "Verbose") // 注册别名（此处loud和Verbose建立了别名）
	viper.Set("verbose", true)             // 结果与下一行相同
	viper.Set("loud", true)                // 结果与前一行相同
	viper.GetBool("loud")                  // true
	viper.GetBool("verbose")               // true

	r.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version": viper.GetString("version"), // 配置文件中的值会覆盖掉默认值
		})
	})

	r.Run(":9090")
}
