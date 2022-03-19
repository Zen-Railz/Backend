package log

import (
	"strings"
	"zenrailz/environment"
)

type LogLevel int

const (
	dbg LogLevel = iota
	inf
	err
)

var level LogLevel

func SetLevel() {
	lvl, err := environment.LogLevel()

	if err != nil {
		level = parseLevel(lvl)
	} else {
		level = inf
	}
}

func parseLevel(value string) LogLevel {
	switch {
	case isDebug(value):
		return dbg
	case isError(value):
		return err
	default:
		return inf
	}
}

func isDebug(value string) bool {
	return strings.EqualFold(value, "dbg") ||
		strings.EqualFold(value, "debug")
}

func isError(value string) bool {
	return strings.EqualFold(value, "err") ||
		strings.EqualFold(value, "error")
}
