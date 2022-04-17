package errorr

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
