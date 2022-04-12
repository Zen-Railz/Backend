package mock

import "zenrailz/errorr"

func NewError() *Error {
	return &Error{}
}

type Error struct {
	code              string
	displayMessage    string
	stackTraceMessage string
	equalityOutcome   bool
}

func (e *Error) Code() string {
	return e.code
}

func (e *Error) Error() string {
	return e.displayMessage
}

func (e *Error) Is(target error) bool {
	return e.equalityOutcome
}

func (e *Error) Elaborate() string {
	return e.stackTraceMessage
}

func (e *Error) Trace() errorr.Entity {
	return e
}

func (e *Error) SetCode(code string) *Error {
	e.code = code
	return e
}

func (e *Error) SetDisplayMessage(message string) *Error {
	e.displayMessage = message
	return e
}

func (e *Error) SetStackTraceMessage(message string) *Error {
	e.stackTraceMessage = message
	return e
}

func (e *Error) SetEquality(isEqual bool) *Error {
	e.equalityOutcome = isEqual
	return e
}
