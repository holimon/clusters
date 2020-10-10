// +build mipsle mips mips64le mips64

package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var AppLogger, _ = zap.NewProduction()

func LoggerInit() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "AppLogger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                   // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), // 打印到控制台
		atomicLevel, // 日志级别
	)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	//开启文件及行号
	development := zap.Development()
	AppLogger = zap.New(core, caller, development)
	AppLogger = zap.New(core)
	return
}
