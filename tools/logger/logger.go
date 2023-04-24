package logger

import (
	"fmt"
	"runtime"
	"time"
)

type LogType int

const (
	Success LogType = 1
	Warning         = 2
	Error           = 3
	Info            = 4
)

func Log(msg string, logType LogType) {
	programCounter, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(programCounter)
	fmt.Printf("[%s:%s %d] %s %s: %s\n", file, fn.Name(), line, time.Now().Format("2006-01-02 15:04:05"), getLogTypeHeader(logType), msg)
}

func LogErr(err error) {

	// todo: hack to reduce amount of logging
	if err.Error() == "not enough free memory" {
		return
	}

	if err == nil {
		return
	}
	_, fileName, lineNumber, _ := runtime.Caller(1)

	fmt.Printf("%s %s: %s\n at: %s:%d \n",
		time.Now().Format("2006-01-02 15:04:05"),
		getLogTypeHeader(Error),
		err.Error(),
		fileName,
		lineNumber)
}

func getLogTypeHeader(t LogType) string {
	switch t {
	case Success:
		return "[SUCCESS]"
	case Warning:
		return "[WARNING]"
	case Error:
		return "[ERROR]"
	case Info:
		return "[INFO]"
	}
	return ""
}
