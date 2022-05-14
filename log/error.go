package log

import (
	"fmt"
	gologger "log"
	"runtime/debug"
)

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
