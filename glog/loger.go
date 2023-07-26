package gloger

import (
	"Gerver/giface"
	"fmt"
	"sync"
	"time"
)

type LogLevel int8

const (
	LogDebug LogLevel = iota
	LogInfo  LogLevel = iota
	LogWarn  LogLevel = iota
	LogError LogLevel = iota
	LogPanic LogLevel = iota
	LogFatal LogLevel = iota
)

var levels = []string{
	"[ DEBUG ]",
	"[ INFO  ]",
	"[ WARN  ]",
	"[ ERROR ]",
	"[ PANIC ]",
	"[ FATAL ]",
}

type Logger struct {
	mu    sync.Mutex
	level LogLevel
	time  time.Time
}

func NewLogger(level LogLevel) giface.ILogger {
	l := &Logger{
		level: level,
	}

	return l
}

func (l *Logger) getTime() []byte {
	t := time.Now()
	year, month, day := t.Date()
	buf := make([]byte, 0)
	buf = append(buf, '<')
	buf = append(buf, fmt.Sprintf("%d", year)...)
	buf = append(buf, '/')
	buf = append(buf, fmt.Sprintf("%d", month)...)
	buf = append(buf, '/')
	buf = append(buf, fmt.Sprintf("%d", day)...)
	buf = append(buf, ' ')

	hour, min, sec := t.Clock()
	buf = append(buf, fmt.Sprintf("%d", hour)...)
	buf = append(buf, ':')
	buf = append(buf, fmt.Sprintf("%d", min)...)
	buf = append(buf, ':')
	buf = append(buf, fmt.Sprintf("%d", sec)...)
	buf = append(buf, '>')
	return buf
}

func (l *Logger) Debug() {
	fmt.Println(levels[0], string(l.getTime()))
}
