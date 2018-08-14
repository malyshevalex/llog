package llog

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var levelPrefix = []string{
	LFatal:   "FF",
	LError:   "EE",
	LWarning: "WW",
	LInfo:    "II",
	LDebug:   "DD",
}

type message struct {
	level Level
	time  time.Time
	str   string
}

type asyncLogger struct {
	level    Level
	out      io.Writer
	in       chan message
	buf      []byte
	flushMtx sync.Mutex
}

const messageBufferSize = 100
const formatBufferSize = 1024

// New creates new Logger object
func New(level Level, out io.Writer) Logger {
	logger := &asyncLogger{
		level: level,
		out:   out,
		in:    make(chan message, messageBufferSize),
		buf:   make([]byte, 0, formatBufferSize),
	}
	go func(l *asyncLogger) {
	LOOP:
		for {
			l.flushMtx.Lock()
			select {
			case msg, ok := <-l.in:
				if ok {
					l.doOutput(&msg)
				} else {
					l.flushMtx.Unlock()
					break LOOP
				}
			}
			l.flushMtx.Unlock()
		}
	}(logger)
	return logger
}

func (l *asyncLogger) flush() {
	l.flushMtx.Lock()
	defer l.flushMtx.Unlock()
	for len(l.in) > 0 {
		msg, ok := <-l.in
		if ok {
			l.doOutput(&msg)
		} else {
			return
		}
	}
}

func (l *asyncLogger) Close() {
	close(l.in)
	l.flush()
}

func (l *asyncLogger) AddPrefix(prefix string) Logger {
	return newPrefix(l, l.level, prefix)
}

func (l *asyncLogger) Fatal(v ...interface{}) {
	l.output(LFatal, fmt.Sprint(v...))
	l.Close()
	os.Exit(1)
}

func (l *asyncLogger) Fatalf(f string, v ...interface{}) {
	l.output(LFatal, fmt.Sprintf(f, v...))
	l.Close()
	os.Exit(1)
}

func (l *asyncLogger) Error(v ...interface{}) {
	l.output(LError, fmt.Sprint(v...))
}

func (l *asyncLogger) Errorf(f string, v ...interface{}) {
	l.output(LError, fmt.Sprintf(f, v...))
}

func (l *asyncLogger) Warning(v ...interface{}) {
	if l.level >= LWarning {
		l.output(LWarning, fmt.Sprint(v...))
	}
}

func (l *asyncLogger) Warningf(f string, v ...interface{}) {
	if l.level >= LWarning {
		l.output(LWarning, fmt.Sprintf(f, v...))
	}
}

func (l *asyncLogger) Info(v ...interface{}) {
	if l.level >= LInfo {
		l.output(LInfo, fmt.Sprint(v...))
	}
}

func (l *asyncLogger) Infof(f string, v ...interface{}) {
	if l.level >= LInfo {
		l.output(LInfo, fmt.Sprintf(f, v...))
	}
}

func (l *asyncLogger) Debug(v ...interface{}) {
	if l.level >= LDebug {
		l.output(LDebug, fmt.Sprint(v...))
	}
}

func (l *asyncLogger) Debugf(f string, v ...interface{}) {
	if l.level >= LDebug {
		l.output(LDebug, fmt.Sprintf(f, v...))
	}
}

func (l *asyncLogger) output(level Level, msg string) {
	l.in <- message{
		level,
		time.Now(),
		msg,
	}
}

func (l *asyncLogger) doOutput(msg *message) {
	buf := l.buf[:0]
	formatHeader(&buf, msg.level, msg.time)
	buf = append(buf, msg.str...)
	if len(msg.str) == 0 || msg.str[len(msg.str)-1] != '\n' {
		buf = append(buf, '\n')
	}
	l.out.Write(buf)
}
