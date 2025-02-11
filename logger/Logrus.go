//  Copyright DataWiseHQ/grule-rule-engine Authors
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
	"github.com/sirupsen/logrus"
)

type logrusLogEntry struct {
	entry *logrus.Entry
}

type logrusLogger struct {
	logger *logrus.Logger
}

func NewLogrus(logger *logrus.Logger) LogEntry {
	l := logrusLogger{logger: logger}

	return l.WithFields(Fields{"lib": "grule-rule-engine"})
}

func (l *logrusLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *logrusLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *logrusLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *logrusLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *logrusLogger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *logrusLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *logrusLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *logrusLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *logrusLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l *logrusLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l *logrusLogger) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

func (l *logrusLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

func (l *logrusLogger) WithFields(fields Fields) LogEntry {

	return LogEntry{
		Logger: &logrusLogEntry{
			entry: l.logger.WithFields(convertToLogrusFields(fields)),
		},
		Level: convertLogrusToInternalLevel(l.logger.GetLevel()),
	}
}

func (l *logrusLogEntry) Trace(args ...interface{}) {
	l.entry.Trace(args...)
}

func (l *logrusLogEntry) Tracef(format string, args ...interface{}) {
	l.entry.Tracef(format, args...)
}

func (l *logrusLogEntry) Print(args ...interface{}) {
	l.entry.Info(args...)
}

func (l *logrusLogEntry) Println(args ...interface{}) {
	l.entry.Infoln(args...)
}

func (l *logrusLogEntry) Printf(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

func (l *logrusLogEntry) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

func (l *logrusLogEntry) Info(args ...interface{}) {
	l.entry.Info(args...)
}

func (l *logrusLogEntry) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

func (l *logrusLogEntry) Error(args ...interface{}) {
	l.entry.Error(args...)
}

func (l *logrusLogEntry) Panic(args ...interface{}) {
	l.entry.Panic(args...)
}

func (l *logrusLogEntry) Fatal(args ...interface{}) {
	l.entry.Fatal(args...)
}

func (l *logrusLogEntry) Debugf(template string, args ...interface{}) {
	l.entry.Debugf(template, args...)
}

func (l *logrusLogEntry) Infof(template string, args ...interface{}) {
	l.entry.Infof(template, args...)
}

func (l *logrusLogEntry) Warnf(template string, args ...interface{}) {
	l.entry.Warnf(template, args...)
}

func (l *logrusLogEntry) Errorf(template string, args ...interface{}) {
	l.entry.Errorf(template, args...)
}

func (l *logrusLogEntry) Panicf(template string, args ...interface{}) {
	l.entry.Panicf(template, args...)
}

func (l *logrusLogEntry) Fatalf(template string, args ...interface{}) {
	l.entry.Fatalf(template, args...)
}

func (l *logrusLogEntry) WithFields(fields Fields) LogEntry {

	return LogEntry{
		Logger: &logrusLogEntry{
			entry: l.entry.WithFields(convertToLogrusFields(fields)),
		},
		Level: convertLogrusToInternalLevel(l.entry.Level),
	}
}

func (l *logrusLogger) Trace(args ...interface{}) {
	l.logger.Trace(args...)
}

func (l *logrusLogger) Tracef(format string, args ...interface{}) {
	l.logger.Tracef(format, args...)
}

func (l *logrusLogger) Print(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *logrusLogger) Println(args ...interface{}) {
	l.logger.Infoln(args...)
}

func (l *logrusLogger) Printf(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func convertToLogrusFields(fields Fields) logrus.Fields {
	logrusFields := logrus.Fields{}

	for index, val := range fields {
		logrusFields[index] = val
	}

	return logrusFields
}

func convertLogrusToInternalLevel(level logrus.Level) Level {
	switch level {
	case logrus.TraceLevel:

		return TraceLevel
	case logrus.DebugLevel:

		return DebugLevel
	case logrus.InfoLevel:

		return InfoLevel
	case logrus.WarnLevel:

		return WarnLevel
	case logrus.ErrorLevel:

		return ErrorLevel
	case logrus.FatalLevel:

		return FatalLevel
	case logrus.PanicLevel:

		return PanicLevel
	default:

		return DebugLevel
	}
}
