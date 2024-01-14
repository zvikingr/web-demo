package log

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

// Config logger config
type Config struct {
	LogFile    string `toml:"log_file"`    // 日志文件路径
	Level      string `toml:"level"`       // 最低日志级别
	MaxSize    int    `toml:"max_size"`    // 单个日志最大容量，单位：兆
	MaxBackups int    `toml:"max_backups"` // 最多保存多少个归档
	MaxAge     int    `toml:"max_age"`     // 最长保存时间，单位：小时
	Compress   bool   `toml:"compress"`    // 是否对归档文件进行压缩
}

// InitLog init global logger
func InitLog(cfg *Config) {
	var (
		encoder       = getEncoder()       // 控制日志输出样式
		fileSyncer    = getFileSyncer(cfg) // 输出到文件
		consoleSyncer = getConsoleSyncer() // 输出到控制台
	)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, consoleSyncer, covertLevel(cfg.Level)),
		zapcore.NewCore(encoder, fileSyncer, covertLevel(cfg.Level)),
	)

	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getConsoleSyncer() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func getFileSyncer(cfg *Config) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   cfg.LogFile,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func covertLevel(levelStr string) zapcore.Level {
	switch levelStr {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
