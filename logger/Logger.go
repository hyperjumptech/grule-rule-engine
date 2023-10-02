//  Copyright hyperjumptech/grule-rule-engine Authors
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
	"go.uber.org/zap"
)

// Level type
type Level uint32

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

type LogEntry struct {
	Logger
	Level Level
}

type Fields map[string]interface{}

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})

	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})

	Trace(args ...interface{})
	Tracef(format string, args ...interface{})

	Print(args ...interface{})
	Println(args ...interface{})
	Printf(format string, args ...interface{})

	WithFields(keyValues Fields) LogEntry
}

var (
	Log LogEntry
)

func init() {
	logger := logrus.New()
	logger.Level = logrus.InfoLevel

	Log = LogEntry{
		Logger: NewLogrus(logger).WithFields(Fields{"lib": "grule-rule-engine"}),
		Level:  DebugLevel,
	}
}

// SetLogger changes default logger on external
func SetLogger(externalLog interface{}) {
	Log = NewLogEntry(externalLog)
}

// SetLogLevel will set the logger log level
func SetLogLevel(lvl Level) {
	Log.Level = lvl
}

// NewLogEntry creates a LogEntry instance with log, log should be *zap.Logger or *logrus.Logger
func NewLogEntry(log any) (logger LogEntry) {
	switch _log := log.(type) {
	case *zap.Logger:
		logger = NewZap(_log)

	case *logrus.Logger:
		logger = NewLogrus(_log)

	}

	return
}
