package initalize

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"sync"
	"time"
)

var (
	logInstance *zap.Logger
	once sync.Once
)

func Logger() *zap.Logger {
	once.Do(func() {
		logInstance = logger()
	})

	return logInstance
}


func logger()  *zap.Logger{
	//配置logger
	// 正确日志输出到文件
	infoWriter := lumberjack.Logger{
		Filename:   "log/info.log",
	}
	// 错误日志输出到文件
	errorWriter := lumberjack.Logger{
		Filename:    "log/error.log",
	}

	// 日志输出格式
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "",
		MessageKey:     "msg",
		StacktraceKey:  "",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // FullCallerEncoder
		EncodeName:     zapcore.FullNameEncoder,
	}

	consoleConfig := zapcore.EncoderConfig{
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "",
		CallerKey:      "",
		MessageKey:     "msg",
		StacktraceKey:  "",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 根据配置调整日志级别, 支持http接口动态修改zap日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)

	// 设置输出源，输出格式，日志等级
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(&infoWriter), zap.InfoLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(&errorWriter), zap.ErrorLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(consoleConfig), zapcore.AddSync(os.Stdout), atomicLevel),
	)

	// 开启文件及行号
	development := zap.Development()

	// 构造日志
	z := zap.New(core, development)

	z = z.With(zap.Int64("uuid", time.Now().Unix()))

	return z
}