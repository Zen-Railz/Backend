package log

func New() *logger {
	return &logger{}
}

type logger struct{}
