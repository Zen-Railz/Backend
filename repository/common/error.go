package common

import (
	"fmt"
	"zenrailz/errorr"
)

func ParseError(code string, message string, err error) errorr.Entity {
	errMsg := fmt.Sprintf("%s %s", message, err.Error())

	return errorr.New(code, errMsg, nil)
}
