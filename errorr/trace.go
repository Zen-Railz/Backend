package errorr

import "runtime"

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
