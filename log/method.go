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
		note := ""

		if message != "" {
			note += message
		}

		if e != nil {
			note = fmt.Sprintf("%s\n%+v", note, e)
		}

		note = fmt.Sprintf("%s\n%s", note, string(debug.Stack()))

		logger.Println(note)
	}
}
