package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

/*
*
  - Description:
    使用lumberjack库对日志进行切割
  - @Author Hollis
  - @Create 2023/9/28 14:25
*/
var logger05 *zap.Logger

func getLogWriter05() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./src/14_zap/05.lumberjack_log.log",
		MaxSize:    1,     // 最大日志大小，单位为M，默认为100M
		MaxBackups: 5,     // 最大备份数量
		MaxAge:     30,    // 最大备份天数
		Compress:   false, // 是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

func InitLogger05() {
	encoder := func() zapcore.Encoder { // 通过匿名函数，设置编码时的配置信息
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		return zapcore.NewConsoleEncoder(encoderConfig)
	}()

	writeSyncer := getLogWriter05()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger05 = zap.New(core, zap.AddCaller())
}

func main() {
	InitLogger05()
	defer logger05.Sync()

	// 测试日志分割的功能
	for i := 0; i < 20000; i++ {
		logger05.Info("test for log rotate...")
	}
}
