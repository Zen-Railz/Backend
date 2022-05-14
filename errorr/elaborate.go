package errorr

import "fmt"

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
