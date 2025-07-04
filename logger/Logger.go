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
	"github.com/rs/zerolog"
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
	Log = NewNoopLogger()
}

// SetLogger changes default logger on external
//
// logrusLogger := logrus.New()
// SetLogger(logrusLogger)
func SetLogger(externalLog interface{}) {
	switch externalLog.(type) {
	case *zap.Logger:
		log, ok := externalLog.(*zap.Logger)
		if !ok {

			return
		}
		Log = NewZap(log)
	case *logrus.Logger:
		log, ok := externalLog.(*logrus.Logger)
		if !ok {

			return
		}
		Log = NewLogrus(log)
	case *zerolog.Logger:
		log, ok := externalLog.(*zerolog.Logger)
		if !ok {

			return
		}
		Log = NewZero(log)
	default:

		return
	}
}

// SetLogLevel will set the logger log level
func SetLogLevel(lvl Level) {
	Log.Level = lvl
}
