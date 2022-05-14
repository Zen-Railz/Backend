package log

import (
	"fmt"
	gologger "log"
)

func (l *logger) Info(message string) {
	if level <= inf {
		gologger.SetPrefix("[INF] ")
		note := fmt.Sprintf("%v", message)
		gologger.Println(note)
	}
}
