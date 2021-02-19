/*
Copyright © 2021 Markus Mahlberg <markus@mahlberg.io>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mocked

import (
	"io"
	"log"

	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/mock"
)

// Logger is the primary mock object.
type Logger struct {
	mock.Mock
}

// Name satisfies the hclog.Logger interface.
func (m *Logger) Name() string {
	rv := m.Called()
	return rv.String(0)
}

// Named satisfies the hclog.Logger interface.
func (m *Logger) Named(n string) hclog.Logger {
	rv := m.Called()
	return rv.Get(0).(hclog.Logger)
}

// ResetNamed satisfies the hclog.Logger interface.
func (m *Logger) ResetNamed(n string) hclog.Logger {
	rv := m.Called(n)
	return rv.Get(0).(hclog.Logger)
}

// SetLevel satisfies the hclog.Logger interface.
func (m *Logger) SetLevel(lv hclog.Level) {
	m.Called()
}

// StandardLogger satisfies the hclog.Logger interface.
func (m *Logger) StandardLogger(opts *hclog.StandardLoggerOptions) *log.Logger {
	rv := m.Called(opts)
	return rv.Get(0).(*log.Logger)
}

// StandardWriter satisfies the hclog.Logger interface.
func (m *Logger) StandardWriter(opts *hclog.StandardLoggerOptions) io.Writer {
	rv := m.Called(opts)
	return rv.Get(0).(io.Writer)
}

// With satisfies the hclog.Logger interface.
func (m *Logger) With(args ...interface{}) hclog.Logger {
	rv := m.Called(args...)
	return rv.Get(0).(hclog.Logger)
}

// Log satisfies the hclog.Logger interface.
func (m *Logger) Log(level hclog.Level, msg string, args ...interface{}) {
	m.Called()
}

// Trace satisfies the hclog.Logger interface.
func (m *Logger) Trace(msg string, args ...interface{}) {
	m.Called()
}

// IsTrace satisfies the hclog.Logger interface.
func (m *Logger) IsTrace() bool {
	rv := m.Called()
	return rv.Bool(0)
}

// Debug satisfies the hclog.Logger interface.
func (m *Logger) Debug(msg string, args ...interface{}) {
	m.Called()
}

// IsDebug satisfies the hclog.Logger interface.
func (m *Logger) IsDebug() bool {
	rv := m.Called()
	return rv.Bool(0)
}

// Info satisfies the hclog.Logger interface.
func (m *Logger) Info(msg string, args ...interface{}) {
	m.Called()
}

// IsInfo satisfies the hclog.Logger interface.
func (m *Logger) IsInfo() bool {
	rv := m.Called()
	return rv.Bool(0)
}

// Warn satisfies the hclog.Logger interface.
func (m *Logger) Warn(msg string, args ...interface{}) {
	m.Called()
}

// IsWarn satisfies the hclog.Logger interface.
func (m *Logger) IsWarn() bool {
	rv := m.Called()
	return rv.Bool(0)
}

// Error satisfies the hclog.Logger interface.
func (m *Logger) Error(msg string, args ...interface{}) {
	m.Called()
}

// IsError satisfies the hclog.Logger interface.
func (m *Logger) IsError() bool {
	rv := m.Called()
	return rv.Bool(0)
}

// ImpliedArgs satisfies the hclog.Logger interface.
func (m *Logger) ImpliedArgs() []interface{} {
	rv := m.Called()
	return rv.Get(0).([]interface{})
}
