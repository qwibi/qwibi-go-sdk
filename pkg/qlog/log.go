package qlog

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	// . "github.com/logrusorgru/aurora"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
	//log.SetFormatter(&log.JSONFormatter{})
	//log.SetLevel(log.InfoLevel)
}

func caller(pc uintptr, file string, line int, ok bool) string {
	current, _ := os.Getwd()
	rel, _ := filepath.Rel(current, file)
	return fmt.Sprintf("%s%s%d %s", rel, ":", line, ">")
	// return fmt.Sprintf("%s%s%d %s", Bold(rel), Bold(":"), Bold(line), Blue(">"))
}

func toString(msg ...interface{}) string {
	s := []string{}
	for _, m := range msg {
		switch v := m.(type) {
		case error:
			s = append(s, fmt.Sprintf("%v", v.Error()))
		default:
			s = append(s, fmt.Sprintf("%v", m))
		}
	}

	return strings.Join(s, " ")
}

func Error(msg ...interface{}) error {
	if msg == nil {
		return nil
	}
	caller := caller(runtime.Caller(1))
	m := toString(msg...)
	log.Error(caller, " ", m)
	return fmt.Errorf(m)
}

func Errorf(format string, args ...interface{}) error {
	if args == nil {
		return nil
	}
	caller := caller(runtime.Caller(1))
	m := fmt.Sprintf(format, args...)
	log.Error(caller, " ", m)
	return fmt.Errorf(m)
}

func Info(msg ...interface{}) {
	caller := caller(runtime.Caller(1))
	m := toString(msg...)
	log.Info(caller, " ", m)
}

func Infof(format string, args ...interface{}) {
	caller := caller(runtime.Caller(1))
	m := fmt.Sprintf(format, args...)
	log.Info(caller, " ", m)
}

func Warn(msg ...interface{}) {
	caller := caller(runtime.Caller(1))
	m := toString(msg...)
	log.Warn(caller, " ", m)
}

func Warnf(format string, args ...interface{}) {
	caller := caller(runtime.Caller(1))
	m := fmt.Sprintf(format, args...)
	log.Warn(caller, " ", m)
}

func Debug(msg ...interface{}) {
	caller := caller(runtime.Caller(1))
	m := toString(msg...)
	log.Debug(caller, " ", m)
}

func Debugf(format string, args ...interface{}) {
	caller := caller(runtime.Caller(1))
	m := fmt.Sprintf(format, args...)
	log.Debug(caller, " ", m)
}

func Fatal(msg ...interface{}) {
	caller := caller(runtime.Caller(1))
	m := toString(msg...)
	log.Fatal(caller, " ", m)
}

func Fatalf(format string, args ...interface{}) {
	caller := caller(runtime.Caller(1))
	m := fmt.Sprintf(format, args...)
	log.Fatal(caller, " ", m)
}

func TODO(msg ...interface{}) error {
	caller := caller(runtime.Caller(1))
	m := toString(msg...)
	log.Error(caller, " TODO ", m)
	return fmt.Errorf(m)
}

func Status(c codes.Code, msg string) error {
	caller := caller(runtime.Caller(1))
	log.Errorf("%s [%s] %s", caller, c, msg)
	return status.Error(c, msg)
}

func Statusf(c codes.Code, format string, a ...any) error {
	caller := caller(runtime.Caller(1))
	m := toString(a...)
	log.Errorf("%s [%s] %s", caller, c, m)
	return status.Errorf(c, format, a...)
}
