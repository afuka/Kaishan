package log

import (
	"go.uber.org/zap/zapcore"
)

type Field map[string]string

// Fields 使之支持 key => value 的json模式
func Fields(fields Field) (zf []zapcore.Field) {
	for k, v := range fields {
		zf = append(zf, zapcore.Field{
			Key: k,
			String: v,
		})
	}

	return
}