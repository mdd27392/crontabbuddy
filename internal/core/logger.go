package core

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Logger is a tiny leveled logger.
type Logger struct {
	level int
	out *log.Logger
}

const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
)

// NewLogger builds a logger at the given level name.
func NewLogger(level string) *Logger {
	lvl := LevelInfo
	switch strings.ToLower(level) {
	case "debug":
 lvl = LevelDebug
	case "warn", "warning":
 lvl = LevelWarn
	case "error":
 lvl = LevelError
	}
	return &Logger{level: lvl, out: log.New(os.Stderr, "", 0)}
}

func (l *Logger) log(lvl int, tag string, format string, args ...any) {
	if lvl < l.level {
 return
	}
	ts := time.Now().Format("15:04:05")
	msg := fmt.Sprintf(format, args...)
	l.out.Printf("[%s] %-5s %s", ts, tag, msg)
}

func (l *Logger) Debugf(format string, args ...any) { l.log(LevelDebug, "DEBUG", format, args...) }
func (l *Logger) Infof(format string, args ...any) { l.log(LevelInfo, "INFO", format, args...) }
func (l *Logger) Warnf(format string, args ...any) { l.log(LevelWarn, "WARN", format, args...) }
func (l *Logger) Errorf(format string, args ...any) { l.log(LevelError, "ERROR", format, args...) }