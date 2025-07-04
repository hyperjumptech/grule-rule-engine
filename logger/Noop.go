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

type NoopLogger struct{}

func NewNoopLogger() LogEntry {
	noop := &NoopLogger{}
	return noop.WithFields(nil)
}

func (logger *NoopLogger) Debug(args ...interface{}) {}
func (logger *NoopLogger) Info(args ...interface{})  {}
func (logger *NoopLogger) Warn(args ...interface{})  {}
func (logger *NoopLogger) Error(args ...interface{}) {}
func (logger *NoopLogger) Panic(args ...interface{}) {}
func (logger *NoopLogger) Fatal(args ...interface{}) {}

func (logger *NoopLogger) Debugf(template string, args ...interface{}) {}
func (logger *NoopLogger) Infof(template string, args ...interface{})  {}
func (logger *NoopLogger) Warnf(template string, args ...interface{})  {}
func (logger *NoopLogger) Errorf(template string, args ...interface{}) {}
func (logger *NoopLogger) Panicf(template string, args ...interface{}) {}
func (logger *NoopLogger) Fatalf(template string, args ...interface{}) {}

func (logger *NoopLogger) Trace(args ...interface{})                 {}
func (logger *NoopLogger) Tracef(format string, args ...interface{}) {}

func (logger *NoopLogger) Print(args ...interface{})                 {}
func (logger *NoopLogger) Println(args ...interface{})               {}
func (logger *NoopLogger) Printf(format string, args ...interface{}) {}

func (logger *NoopLogger) WithFields(keyValues Fields) LogEntry {
	return LogEntry{
		Logger: logger,
		Level:  TraceLevel,
	}
}
