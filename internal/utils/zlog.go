package utils

import (
	"MyBlogs/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func InitLogger(config *config.Config) {

	//管理日志文件，按时间分片
	writer := &lumberjack.Logger{
		Filename:   config.Log.LogPath,
		MaxSize:    100,
		MaxBackups: 30,
		MaxAge:     30,
		LocalTime:  true,
	}

	//快速配置zap
	cfg := zap.NewDevelopmentConfig()
	//修改时间格式
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	//创建core
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		zapcore.AddSync(writer),
		zapcore.DebugLevel,
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
