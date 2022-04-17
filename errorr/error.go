package errorr

import (
	"fmt"
	"runtime"
)

func New(code string, message string, annex interface{}) *entity {
	return (&entity{
		code:    code,
		message: message,
		annex:   annex,
	}).trace(2)
}

func (e *entity) Code() string {
	return e.code
}

func (e *entity) Error() string {
	return fmt.Sprintf("(%s) %s", e.code, e.message)
}

func (e *entity) Is(target error) bool {
	t, ok := target.(*entity)
	if ok {
		return e.code == t.code
	} else {
		return false
	}
}

func (e *entity) Elaborate() string {
	history := ""
	for i, record := range e.stackTrace {
		history += fmt.Sprintf("%d. %s | Line %d\n", i+1, record.FunctionName, record.LineNumber)
	}

	display := fmt.Sprintf("(%s) %s\n%s", e.code, e.message, history)

	if e.annex != nil {
		display = fmt.Sprintf("%s%+v", display, e.annex)
	}

	return display
}

func (e *entity) Trace() Entity {
	return e.trace(2)
}

func (e *entity) trace(offset int) *entity {
	counter, _, lineNumber, success := runtime.Caller(offset)

	if success {
		caller := runtime.FuncForPC(counter).Name()
		e.stackTrace = append(e.stackTrace, StackTrace{
			FunctionName: caller,
			LineNumber:   lineNumber,
		})
	}

	return e
}
