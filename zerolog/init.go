package zerolog

import (
	"io"
	"time"

	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	File string `yaml:"file" json:"file"`
}

func Init(cfg *Config) {
	var w io.Writer
	if file := cfg.File; file == "" {
		w = zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
			w.Out = colorable.NewColorableStderr()
			w.TimeFormat = time.RFC3339
		})
	} else {
		w = &lumberjack.Logger{
			Filename:   file,
			MaxBackups: 5,
			Compress:   true,
		}
	}
	log.Logger = zerolog.New(w).With().Timestamp().Logger()
}
