package main

import (
	"go.uber.org/zap"
	"net/http"
)

/*
*
  - Description:
    zap Logger的基本使用1：将日志输出到控制台
  - @Author Hollis
  - @Create 2023/9/28 13:30
*/
var logger02 *zap.Logger

func InitLogger02() {
	logger02, _ = zap.NewProduction() // 创建一个Logger
}

func simpleHttpGet02(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger02.Error( // 设置logger的级别为Error。默认情况下，日志都会打印到应用程序的console界面。
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger02.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

func main() {
	InitLogger02()
	defer logger02.Sync() // 退出前调用，将缓冲区的数据写入磁盘
	simpleHttpGet02("www.google.com")
	simpleHttpGet02("https://www.google.com")
}
