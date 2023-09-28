package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
)

/*
*
  - Description:
    zap Logger的基本使用3：记录全量日志的同时，将err日志单独输出到文件
  - @Author Hollis
  - @Create 2023/9/28 14:09
*/
var logger04 *zap.Logger

func InitLogger04() {
	encoder := func() zapcore.Encoder { // 通过匿名函数，设置编码时的配置信息
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		return zapcore.NewConsoleEncoder(encoderConfig)
	}()
	// zap_log.log记录全量日志
	logF, _ := os.OpenFile("./src/14_zap/04.zap_log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	c1 := zapcore.NewCore(encoder, zapcore.AddSync(logF), zapcore.DebugLevel)
	// zap_log.err.log记录ERROR级别的日志
	errF, _ := os.OpenFile("./src/14_zap/04.zap_log.err.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	c2 := zapcore.NewCore(encoder, zapcore.AddSync(errF), zap.ErrorLevel)
	// 使用NewTee将c1和c2合并到core
	core := zapcore.NewTee(c1, c2)
	logger04 = zap.New(core, zap.AddCaller())
}
func simpleHttpGet04(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger04.Error( // 设置logger的级别为Error。默认情况下，日志都会打印到应用程序的console界面。
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger04.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

func main() {
	InitLogger04()
	defer logger04.Sync() // 退出前调用，将缓冲区的数据写入磁盘
	simpleHttpGet04("www.google.com")
	simpleHttpGet04("https://www.google.com")
}
