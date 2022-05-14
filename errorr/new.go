package errorr

func New(code string, message string, annex interface{}) *entity {
	return (&entity{
		code:    code,
		message: message,
		annex:   annex,
	}).trace(2)
}

type entity struct {
	code       string
	message    string
	annex      interface{}
	stackTrace []StackTrace
}

type StackTrace struct {
	FunctionName string
	LineNumber   int
}
