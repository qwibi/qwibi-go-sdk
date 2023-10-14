package qlog

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	// . "github.com/logrusorgru/aurora"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
	//log.SetLevel(log.InfoLevel)
}

type Code int64

const (
	TRACE Code = 1
	DEBUG Code = 2
	INFO  Code = 3
	WARN  Code = 4
	ERROR Code = 5
	FATAL Code = 6
)

type Status struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

func (c Status) Error() string {
	json, _ := json.Marshal(c)
	return string(json)
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
		case Status:
			s = append(s, fmt.Sprintf("%v", v.Message))
		case error:
			s = append(s, fmt.Sprintf("%v", v.Error()))
		default:
			s = append(s, fmt.Sprintf("%v", m))
		}
	}

	return strings.Join(s, " ")
}

func Error(msg ...interface{}) Status {
	caller := caller(runtime.Caller(1))
	m := toString(msg...)

	log.Errorf("%s %s", caller, m)

	return Status{
		Code:    ERROR,
		Message: m,
	}
}

func Errorf(format string, args ...interface{}) Status {
	// err = trace(err)
	caller := caller(runtime.Caller(1))
	m := fmt.Sprintf(format, args...)

	log.Errorf("%s %s", caller, m)

	return Status{
		Code:    ERROR,
		Message: m,
	}
}

func Info(msg ...interface{}) Status {
	caller := caller(runtime.Caller(1))
	m := toString(msg...)

	log.Infof("%s %s", caller, m)

	return Status{
		Code:    INFO,
		Message: m,
	}
}

func Infof(format string, args ...interface{}) Status {
	caller := caller(runtime.Caller(1))
	m := fmt.Sprintf(format, args...)

	log.Infof("%s %d %s", caller, runtime.NumGoroutine(), m)

	return Status{
		Code:    INFO,
		Message: m,
	}
}

func Warn(msg ...interface{}) Status {
	caller := caller(runtime.Caller(1))
	m := toString(msg...)

	log.Warnf("%s %s", caller, m)

	return Status{
		Code:    WARN,
		Message: m,
	}
}

func Warnf(format string, args ...interface{}) Status {
	caller := caller(runtime.Caller(1))
	m := fmt.Sprintf(format, args...)

	log.Warnf("%s %s", caller, m)

	return Status{
		Code:    WARN,
		Message: m,
	}
}

func Debug(msg ...interface{}) Status {
	caller := caller(runtime.Caller(1))
	m := toString(msg...)

	log.Debugf("%s %s", caller, m)

	return Status{
		Code:    DEBUG,
		Message: m,
	}
}

func Debugf(format string, args ...interface{}) Status {
	caller := caller(runtime.Caller(1))
	m := fmt.Sprintf(format, args...)

	log.Debugf("%s %s", caller, m)

	return Status{
		Code:    DEBUG,
		Message: m,
	}
}

func Fatal(msg ...interface{}) Status {
	caller := caller(runtime.Caller(1))
	m := toString(msg...)

	log.Fatalf("%s %s", caller, m)

	return Status{
		Code:    FATAL,
		Message: m,
	}
}

func Fatalf(format string, args ...interface{}) Status {
	caller := caller(runtime.Caller(1))
	m := fmt.Sprintf(format, args...)

	log.Fatalf("%s %s", caller, m)

	return Status{
		Code:    FATAL,
		Message: m,
	}
}

func TODO(msg ...interface{}) Status {
	caller := caller(runtime.Caller(1))
	m := toString(msg...)

	log.Warnf("%s TODO %s", caller, m)

	return Status{
		Code:    WARN,
		Message: m,
	}
}

func BedRequest(msg ...interface{}) Status {
	caller := caller(runtime.Caller(1))
	m := toString(msg...)

	log.Warnf("%s bed request %s", caller, m)

	return Status{
		Code:    WARN,
		Message: "bed request",
	}
}
