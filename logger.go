package immortal

import (
	"io"
	"log"
)

// Logger interface
type Logger interface {
	Log(input io.ReadCloser)
	IsLogging() bool
}

// LogWriter implements Logger
type LogWriter struct {
	logger *log.Logger
}

// Log write to the logger
func (l *LogWriter) Log(input io.ReadCloser) { _ = "STUB: not implemented"; return }

// If non EOF error happens, we dump the log and not to close the pipe.
// Calling input.Close() right away may trigger SIGPIPE signal to the daemon and restart.

// IsLogging return true if an availale logger exists
func (l *LogWriter) IsLogging() bool { _ = "STUB: not implemented"; return false }

// NewStderrLogger return  Logger instance
func NewStderrLogger(cfg *Config) *log.Logger { _ = "STUB: not implemented"; return nil }

// NewLogger return a Logger instance
func NewLogger(cfg *Config, quit chan struct{}) *log.Logger { _ = "STUB: not implemented"; return nil }

// create a multiwriter

// keep logger up

// add writer

// create the logger
