package log

import (
	"fmt"
	gologger "log"
	"runtime/debug"
)

type logger struct{}

func New() *logger {
	return &logger{}
}

func (l *logger) Debug(message string) {
	if level <= dbg {
		gologger.SetPrefix("[DBG] ")
		note := fmt.Sprintf("%v", message)
		gologger.Println(note)
	}
}

func (l *logger) Info(message string) {
	if level <= inf {
		gologger.SetPrefix("[INF] ")
		note := fmt.Sprintf("%v", message)
		gologger.Println(note)
	}
}

func (l *logger) Error(message string, e interface{}) {
	if level <= err {
		gologger.SetPrefix("[ERR] ")
		note := ""

		if message != "" {
			note += message
		}

		if e != nil {
			note = fmt.Sprintf("%s\n%+v", note, e)
		}

		note = fmt.Sprintf("%s\n%s", note, string(debug.Stack()))

		gologger.Println(note)
	}
}
