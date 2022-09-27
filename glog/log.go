package glog

import (
	"fmt"
	"log"
	"os"
)

type Level int

var (
	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Setup() {
	logger = log.New(os.Stdout, "[INFO]", log.LstdFlags)
}

func Debugf(format string, v ...interface{}) {
	setPrefix(DEBUG)
	logger.Printf(format, v...)
}
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

func Infof(format string, v ...interface{}) {
	setPrefix(INFO)
	logger.Printf(format, v...)
}
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v...)
}

func Warnf(format string, v ...interface{}) {
	setPrefix(WARNING)
	logger.Printf(format, v...)
}
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v...)
}

func Errorf(format string, v ...interface{}) {
	setPrefix(ERROR)
	logger.Printf(format, v...)
}
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v...)
}

func Fatalf(format string, v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalf(format, v...)
}
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v...)
}

func setPrefix(level Level) {
	logPrefix = fmt.Sprintf("[%5s] ", levelFlags[level])
	logger.SetPrefix(logPrefix)
}
