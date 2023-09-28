package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"net/http"
	"os"
)

/*
*
  - Description:
    zap Logger的基本使用2：将日志输出日志文件，包括设置日志级别、更改时间编码、添加调用者信息等。
  - @Author Hollis
  - @Create 2023/9/28 13:44
*/
var logger03 *zap.Logger // 由于前面的文件中已经定义了全局变量logger，所以这里无需再次定义

func InitLogger03() {
	writeSyncer := getLogWriter03() // 指定日志将写到哪里去
	encoder := getEncoder()         // 编码器(如何写入日志)。这里使用开箱即用的`NewJSONEncoder()`
	logLevel := zapcore.DebugLevel  // 指定写入的日志级别。DebugLevel、InfoLevel、WarnLevel、ErrorLevel、DPanicLevel、PanicLevel、FatalLevel（依次递增）
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)
	logger03 = zap.New(core, zap.AddCaller()) // 通过New()创建一个自定义的Logger。zap.AddCaller()可以在日志中添加调用者信息
}

func getEncoder() zapcore.Encoder { // 设置日志文件的编码方式
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // 时间是以人类可读的方式展示
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 使用大写字母记录日志级别
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())  // Json格式的log
	return zapcore.NewConsoleEncoder(encoderConfig) // 普通格式的log
}

func getLogWriter03() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./src/14_zap/03.zap_log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	ws := io.MultiWriter(file, os.Stdout) // 利用io.MultiWriter支持文件和终端两个输出目标
	return zapcore.AddSync(ws)
}
func simpleHttpGet03(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger03.Error( // 设置logger的级别为Error。默认情况下，日志都会打印到应用程序的console界面。
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger03.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

func main() {
	InitLogger03()
	defer logger03.Sync() // 退出前调用，将缓冲区的数据写入磁盘
	simpleHttpGet03("www.google.com")
	simpleHttpGet03("https://www.google.com")
}
