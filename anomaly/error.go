package anomaly

import (
	"fmt"
	"runtime"
)

type ServiceError struct {
	Code    string
	Message string
	Annex   interface{}
	history []errorHistory
}

type errorHistory struct {
	functionName string
	lineNumber   int
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("(%s) %s", e.Code, e.Message)
}

func (e *ServiceError) Is(target error) bool {
	t, ok := target.(*ServiceError)
	if ok {
		return e.Code == t.Code
	} else {
		return false
	}
}

func (e *ServiceError) Trace() *ServiceError {
	counter, _, lineNumber, success := runtime.Caller(1)

	if success {
		caller := runtime.FuncForPC(counter).Name()
		e.history = append(e.history, errorHistory{
			functionName: caller,
			lineNumber:   lineNumber,
		})
	}

	return e
}

func (e *ServiceError) Elaborate() string {
	history := ""
	for i, record := range e.history {
		history += fmt.Sprintf("%d. %s | Line %d\n", i+1, record.functionName, record.lineNumber)
	}

	display := fmt.Sprintf("(%s) %s\n%s", e.Code, e.Message, history)

	if e.Annex != nil {
		display = fmt.Sprintf("%s%+v", display, e.Annex)
	}

	return display
}
