package environment

import (
	"fmt"
	"os"
	"zenrailz/anomaly"
	"zenrailz/code"
)

func getValue(key string) (string, *anomaly.ServiceError) {
	value, found := os.LookupEnv(key)

	if !found {
		err := &anomaly.ServiceError{
			Code:    code.EnvironmentVariableNotFound,
			Message: fmt.Sprintf("%s not found", key),
		}
		return "", err.Trace()
	}

	return value, nil
}
