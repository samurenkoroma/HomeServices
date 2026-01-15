package logger

import (
	"os"
	"samurenkoroma/services/configs"

	"github.com/rs/zerolog"
)

func NewLogger(cfg configs.LoggerConfig) *zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.Level(cfg.Level))
	var logger zerolog.Logger

	if cfg.Format == "json" {
		logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	} else {
		consoleWriter := zerolog.ConsoleWriter{
			Out: os.Stderr,
		}
		logger = zerolog.New(consoleWriter).With().Timestamp().Logger()
	}

	return &logger
}
