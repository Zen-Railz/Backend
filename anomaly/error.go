package anomaly

import (
	"fmt"
)

type ServiceError struct {
	Code    string
	Message string
	Content interface{}
}

func (e *ServiceError) Error() string {
	if e.Content == nil {
		return fmt.Sprintf("(%s) %v", e.Code, e.Message)
	} else {
		return fmt.Sprintf("(%s) %v\n%+v", e.Code, e.Message, e.Content)
	}
}

func (e *ServiceError) Is(target error) bool {
	t, ok := target.(*ServiceError)
	if ok {
		return e.Code == t.Code
	} else {
		return false
	}
}
