package logger

import (
	"MyBlogs/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func InitLogger(cfg *config.LogConfig) {

	//管理日志文件，按时间分片
	writer := &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}

	// 日志级别
	level := zapcore.InfoLevel
	switch cfg.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	}

	//快速配置zap
	zapCfg := zap.NewDevelopmentConfig()
	//修改时间格式
	zapCfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	//创建core
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zapCfg.EncoderConfig),
		zapcore.AddSync(writer),
		level,
	)

	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func Error(message string, fields ...zap.Field) {
	logger.Error(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	logger.Debug(message, fields...)
}

func Info(message string, fields ...zap.Field) {
	logger.Info(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	logger.Warn(message, fields...)
}
