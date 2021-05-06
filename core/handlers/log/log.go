package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Trace *zap.SugaredLogger // 调试日志
	Info *zap.SugaredLogger // 公共执行日志
	Err *zap.SugaredLogger // 错误日志
)

// InitLogger 初始化日志
func InitLogger() {
	encoder := getEncoder()

	writeSyncerTrace := getLogWriter("./logs/trace.log")
	writeSyncerInfo := getLogWriter("./logs/info.log")
	writeSyncerErr := getLogWriter("./logs/err.log")
	coreTrace := zapcore.NewCore(encoder, writeSyncerTrace, zapcore.DebugLevel)
	coreInfo := zapcore.NewCore(encoder, writeSyncerInfo, zapcore.InfoLevel)
	coreErr := zapcore.NewCore(encoder, writeSyncerErr, zapcore.ErrorLevel)

	logger := zap.New(coreTrace, zap.AddCaller())
	Trace = logger.Sugar()

	logger = zap.New(coreInfo, zap.AddCaller())
	Info = logger.Sugar()

	logger = zap.New(coreErr, zap.AddCaller())
	Err = logger.Sugar()
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