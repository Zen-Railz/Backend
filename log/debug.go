package log

import (
	"fmt"
	gologger "log"
)

func (l *logger) Debug(message string) {
	if level <= dbg {
		gologger.SetPrefix("[DBG] ")
		note := fmt.Sprintf("%v", message)
		gologger.Println(note)
	}
}
