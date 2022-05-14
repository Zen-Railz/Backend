package mock

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Debug(message string)                {}
func (l *Logger) Info(message string)                 {}
func (l *Logger) Error(message string, e interface{}) {}

type Logger struct{}
