package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
)

var logger *zap.Logger

// InitLogger 初始化日志
func InitLogger(dir string) {
	var coreArr []zapcore.Core

	// Encoder
	encoder := getEncoder()

	//日志级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool{  //error级别
		return lev >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {  //info级别
		return lev < zap.ErrorLevel && lev > zap.DebugLevel
	})
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {  //debug级别
		return lev <= zap.DebugLevel
	})

	trace := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(getLogWriter(path.Join(dir, "debug.log")),zapcore.AddSync(os.Stdout)), debugPriority)
	info := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(getLogWriter(path.Join(dir, "info.log")),zapcore.AddSync(os.Stdout)), lowPriority)
	err := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(getLogWriter(path.Join(dir, "err.log")),zapcore.AddSync(os.Stdout)), highPriority)
	coreArr = append(coreArr, trace, info, err)

	logger = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller()) //zap.AddCaller()为显示文件名和行号，可省略
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 写入文件，支持切割日志
func getLogWriter(path string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    10, // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 5,  // 保留旧文件的最大个数
		MaxAge:     30, // 保留旧文件的最大天数
		Compress:   false, // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}