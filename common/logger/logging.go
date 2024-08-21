package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// ANSI 转义码
const (
	ColorReset = "\033[0m"
	ColorRed   = "\033[31m"
	ColorGreen = "\033[32m"
)

var logger *zap.Logger

// CustomColorEncoder 自定义颜色编码器
func CustomColorEncoder(level zapcore.Level, msg string) string {
	switch level {
	case zapcore.InfoLevel:
		return fmt.Sprintf("%s%s%s", ColorGreen, msg, ColorReset)
	case zapcore.ErrorLevel:
		return fmt.Sprintf("%s%s%s", ColorRed, msg, ColorReset)
	default:
		return msg
	}
}

// CustomConsoleEncoder 配置带颜色的 Console Encoder
func CustomConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// InitLogger 初始化全局日志系统，带颜色支持
func InitLogger() {
	// 创建一个自定义的 encoder
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey: "msg",
		LevelKey:   "level",
		TimeKey:    "ts",
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeLevel: func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(CustomColorEncoder(level, level.CapitalString()))
		},
		EncodeCaller: zapcore.ShortCallerEncoder,
	})
	core := zapcore.NewCore(
		encoder,
		zapcore.Lock(os.Stdout),
		zap.DebugLevel,
	)
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	// 将 encoder、输出目标和日志级别组合成 core
	//core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

// Logger returns the initialized logger
//func Logger() *zap.Logger {
//	return logger
//}
