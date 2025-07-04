package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"go.uber.org/zap"
)

type zeroLogger struct {
	zLogger *zerolog.Logger
}

func NewZaro(logger *zerolog.Logger) LogEntry {
	l := zeroLogger{zLogger: logger}

	return l.WithFields(Fields{"lib": "grule-rule-engine"})
}

func (l *zeroLogger) Print(args ...interface{}) {
	l.zLogger.Print(args...)
}

func (l *zeroLogger) Println(args ...interface{}) {
	l.zLogger.Println(args...)
}

func (l *zeroLogger) Trace(args ...interface{}) {
	l.zLogger.Trace().Msg(fmt.Sprint(args...))
}

func (l *zeroLogger) Debug(args ...interface{}) {
	l.zLogger.Trace().Msg(fmt.Sprint(args...))
}

func (l *zeroLogger) Info(args ...interface{}) {
	l.zLogger.Info().Msg(fmt.Sprint(args...))
}

func (l *zeroLogger) Warn(args ...interface{}) {
	l.zLogger.Warn().Msg(fmt.Sprint(args...))
}

func (l *zeroLogger) Error(args ...interface{}) {
	l.zLogger.Error().Msg(fmt.Sprint(args...))
}

func (l *zeroLogger) Panic(args ...interface{}) {
	l.zLogger.Panic().Msg(fmt.Sprint(args...))
}

func (l *zeroLogger) Fatal(args ...interface{}) {
	l.zLogger.Fatal().Msg(fmt.Sprint(args...))
}

func (l *zeroLogger) Printf(template string, args ...interface{}) {
	l.zLogger.Printf(template, args...)
}

func (l *zeroLogger) Tracef(template string, args ...interface{}) {
	l.zLogger.Trace().Msgf(template, args...)
}

func (l *zeroLogger) Debugf(template string, args ...interface{}) {
	l.zLogger.Debug().Msgf(template, args...)
}

func (l *zeroLogger) Infof(template string, args ...interface{}) {
	l.zLogger.Info().Msgf(template, args...)
}

func (l *zeroLogger) Warnf(template string, args ...interface{}) {
	l.zLogger.Warn().Msgf(template, args...)
}

func (l *zeroLogger) Errorf(template string, args ...interface{}) {
	l.zLogger.Error().Msgf(template, args...)
}

func (l *zeroLogger) Panicf(template string, args ...interface{}) {
	l.zLogger.Panic().Msgf(template, args...)
}

func (l *zeroLogger) Fatalf(template string, args ...interface{}) {
	l.zLogger.Fatal().Msgf(template, args...)
}

func (l *zeroLogger) WithFields(fields Fields) LogEntry {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.zLogger.With(f...)

	level := GetLevel(l.sugaredLogger.Desugar().Core())

	return LogEntry{
		Logger: &zapLogger{newLogger},
		Level:  convertZapToInternalLevel(level),
	}
}
