package logger

import (
	"fmt"
	"github.com/rs/zerolog"
)

var _ Logger = (*zeroLogger)(nil)

type zeroLogger struct {
	zLogger *zerolog.Logger
}

func NewZero(logger *zerolog.Logger) LogEntry {
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
	context := l.zLogger.With()

	for k, v := range fields {
		context = context.Interface(k, v)
	}

	newLogger := context.Logger()

	return LogEntry{
		Logger: &zeroLogger{
			zLogger: &newLogger,
		},
		Level: convertZeroLogToInternalLevel(newLogger.GetLevel()),
	}
}

func convertZeroLogToInternalLevel(level zerolog.Level) Level {
	switch level {
	case zerolog.TraceLevel:

		return TraceLevel
	case zerolog.DebugLevel:

		return DebugLevel
	case zerolog.InfoLevel:

		return InfoLevel
	case zerolog.WarnLevel:

		return WarnLevel
	case zerolog.ErrorLevel:

		return ErrorLevel
	case zerolog.FatalLevel:

		return FatalLevel
	case zerolog.PanicLevel:
		return PanicLevel
	default:

		return DebugLevel
	}
}
