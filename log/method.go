package log

import (
	"fmt"
	logger "log"
	"runtime/debug"
)

func Debug(message string) {
	if level <= dbg {
		logger.SetPrefix("[DBG] ")
		note := fmt.Sprintf("%v", message)
		logger.Println(note)
	}
}

func Info(message string) {
	if level <= inf {
		logger.SetPrefix("[INF] ")
		note := fmt.Sprintf("%v", message)
		logger.Println(note)
	}
}

func Error(message string, e interface{}) {
	if level <= err {
		logger.SetPrefix("[ERR] ")
		note := fmt.Sprintf("%v %+v\n%v", message, e, string(debug.Stack()))
		logger.Println(note)
	}
}
