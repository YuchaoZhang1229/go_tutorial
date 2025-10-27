package logger

import (
	"github.com/go-tutorial/08go-cli/config"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init(cfg config.LogConfig) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// 设置日志级别
	level, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)

	// 设置日志格式
	if cfg.Format == "console" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
}

func GetLogger() zerolog.Logger {
	return log.Logger
}
