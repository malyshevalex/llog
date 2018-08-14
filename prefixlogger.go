package log

import (
	"fmt"
	"strings"
)

type prefixLogger struct {
	root   Logger
	prefix string
	level  Level
}

func newPrefix(root Logger, level Level, prefix string) Logger {
	return &prefixLogger{
		root,
		"(" + prefix + ") ",
		level,
	}
}

func (l *prefixLogger) AddPrefix(prefix string) Logger {
	return &prefixLogger{
		l.root,
		strings.TrimSuffix(l.prefix, ") ") + "→" + prefix + ") ",
		l.level,
	}
}

func (l *prefixLogger) Close() {
	// do nothing (it's a virtual logger)
}

// Fatal outputs message with LFatal level and exit 1
func (l *prefixLogger) Fatal(v ...interface{}) {
	l.root.Fatal(l.prefix, fmt.Sprint(v...))
}

func (l *prefixLogger) Fatalf(f string, v ...interface{}) {
	l.root.Fatal(l.prefix, fmt.Sprintf(f, v...))
}

func (l *prefixLogger) Error(v ...interface{}) {
	l.root.Error(l.prefix, fmt.Sprint(v...))
}

func (l *prefixLogger) Errorf(f string, v ...interface{}) {
	l.root.Error(l.prefix, fmt.Sprintf(f, v...))
}

func (l *prefixLogger) Warning(v ...interface{}) {
	if l.level >= LWarning {
		l.root.Warning(l.prefix, fmt.Sprint(v...))
	}
}

func (l *prefixLogger) Warningf(f string, v ...interface{}) {
	if l.level >= LWarning {
		l.root.Warning(l.prefix, fmt.Sprintf(f, v...))
	}
}

func (l *prefixLogger) Info(v ...interface{}) {
	if l.level >= LInfo {
		l.root.Info(l.prefix, fmt.Sprint(v...))
	}
}

func (l *prefixLogger) Infof(f string, v ...interface{}) {
	if l.level >= LInfo {
		l.root.Info(l.prefix, fmt.Sprintf(f, v...))
	}
}

func (l *prefixLogger) Debug(v ...interface{}) {
	if l.level >= LDebug {
		l.root.Debug(l.prefix, fmt.Sprint(v...))
	}
}

func (l *prefixLogger) Debugf(f string, v ...interface{}) {
	if l.level >= LDebug {
		l.root.Debug(l.prefix, fmt.Sprintf(f, v...))
	}
}
