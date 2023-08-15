//  Copyright kalyan-arepalle/grule-rule-engine Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

func NewZap(logger *zap.Logger) LogEntry {
	sugaredLogger := logger.WithOptions(zap.AddCallerSkip(1)).Sugar()
	l := zapLogger{sugaredLogger: sugaredLogger}
	return l.WithFields(Fields{"lib": "grule-rule-engine"})
}

func (l *zapLogger) Print(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

func (l *zapLogger) Println(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	l.sugaredLogger.Info(msg[:len(msg)-1])
}

func (l *zapLogger) Trace(args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

func (l *zapLogger) Debug(args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

func (l *zapLogger) Info(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

func (l *zapLogger) Warn(args ...interface{}) {
	l.sugaredLogger.Warn(args...)
}

func (l *zapLogger) Error(args ...interface{}) {
	l.sugaredLogger.Error(args...)
}

func (l *zapLogger) Panic(args ...interface{}) {
	l.sugaredLogger.Panic(args...)
}

func (l *zapLogger) Fatal(args ...interface{}) {
	l.sugaredLogger.Fatal(args...)
}

func (l *zapLogger) Printf(template string, args ...interface{}) {
	l.sugaredLogger.Infof(template, args...)
}

func (l *zapLogger) Tracef(template string, args ...interface{}) {
	l.sugaredLogger.Debugf(template, args...)
}

func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.sugaredLogger.Debugf(template, args...)
}

func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.sugaredLogger.Infof(template, args...)
}

func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.sugaredLogger.Warnf(template, args...)
}

func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.sugaredLogger.Errorf(template, args...)
}

func (l *zapLogger) Panicf(template string, args ...interface{}) {
	l.sugaredLogger.Panicf(template, args...)
}

func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	l.sugaredLogger.Fatalf(template, args...)
}

func (l *zapLogger) WithFields(fields Fields) LogEntry {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.sugaredLogger.With(f...)

	level := GetLevel(l.sugaredLogger.Desugar().Core())
	return LogEntry{
		Logger: &zapLogger{newLogger},
		Level:  convertZapToInternalLevel(level),
	}
}
func GetLevel(core zapcore.Core) zapcore.Level {
	if core.Enabled(zapcore.DebugLevel) {
		return zapcore.DebugLevel
	}
	if core.Enabled(zapcore.InfoLevel) {
		return zapcore.InfoLevel
	}
	if core.Enabled(zapcore.WarnLevel) {
		return zapcore.WarnLevel
	}
	if core.Enabled(zapcore.ErrorLevel) {
		return zapcore.ErrorLevel
	}
	if core.Enabled(zapcore.DPanicLevel) {
		return zapcore.DPanicLevel
	}
	if core.Enabled(zapcore.PanicLevel) {
		return zapcore.PanicLevel
	}
	if core.Enabled(zapcore.FatalLevel) {
		return zapcore.FatalLevel
	}
	return zapcore.DebugLevel
}

func convertZapToInternalLevel(level zapcore.Level) Level {
	switch level {
	case zapcore.DebugLevel:
		return DebugLevel
	case zapcore.InfoLevel:
		return InfoLevel
	case zapcore.WarnLevel:
		return WarnLevel
	case zapcore.ErrorLevel:
		return ErrorLevel
	case zapcore.DPanicLevel:
		return PanicLevel
	case zapcore.PanicLevel:
		return PanicLevel
	case zapcore.FatalLevel:
		return FatalLevel
	default:
		return DebugLevel
	}
}
