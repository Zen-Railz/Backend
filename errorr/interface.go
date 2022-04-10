package errorr

type Entity interface {
	Code() string
	Error() string
	Is(target error) bool
	Elaborate() string
	Trace() Entity
}
