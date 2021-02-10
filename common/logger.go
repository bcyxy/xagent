package common

import (
	"fmt"
	"path"
	"xagent/glbval"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//ZLog zap logger
var ZLog *zap.Logger

//InitLogger 初始化日志模块
func InitLogger() error {
	logPath := path.Join(glbval.RootPath, "log/xagent.log")

	// 日志分割
	hook := lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    100, //MB
		MaxBackups: 10,
	}
	write := zapcore.AddSync(&hook)

	enConfig := zapcore.EncoderConfig{
		TimeKey:     "time",
		LevelKey:    "level",
		MessageKey:  "msg",
		LineEnding:  zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime:  zapcore.ISO8601TimeEncoder,
	}

	// 设置日志级别
	level := zap.DebugLevel
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(enConfig),
		write,
		level,
	)

	// 构造日志
	ZLog = zap.New(core)
	return nil
}

// LogDebug xxx
func LogDebug(format string, a ...interface{}) {
	logMsg := fmt.Sprintf(format, a...)
	ZLog.Debug(logMsg)
	fmt.Printf("DEBG: %s\n", logMsg)
}

// LogInfo xxx
func LogInfo(format string, a ...interface{}) {
	logMsg := fmt.Sprintf(format, a...)
	ZLog.Info(logMsg)
	fmt.Printf("INFO: %s\n", logMsg)
}

// LogWarn xxx
func LogWarn(format string, a ...interface{}) {
	logMsg := fmt.Sprintf(format, a...)
	ZLog.Warn(logMsg)
	fmt.Printf("WARN: %s\n", logMsg)
}
