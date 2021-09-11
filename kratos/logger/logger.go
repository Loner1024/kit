package logger

import (
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type logger zerolog.Logger

func NewLogger() logger {
	return logger{}
}

// Log implement Kratos log interface
func (l logger) Log(level klog.Level, keyvals ...interface{}) error {
	switch level {
	case klog.LevelDebug:
		log.Debug().Interface("", keyvals).Send()
	case klog.LevelInfo:
		log.Info().Interface("", keyvals).Send()
	case klog.LevelWarn:
		log.Warn().Msgf("%v", keyvals...)
	case klog.LevelError:
		log.Error().Msgf("%v", keyvals...)
	}
	return nil
}
