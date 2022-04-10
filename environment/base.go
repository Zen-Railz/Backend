package environment

import (
	"fmt"
	"os"
	"zenrailz/code"
	"zenrailz/errorr"
)

func getValue(key string) (string, errorr.Entity) {
	value, found := os.LookupEnv(key)

	if !found {
		err := errorr.New(
			code.EnvironmentVariableNotFound,
			fmt.Sprintf("%s not found", key),
			nil,
		)
		return "", err
	}

	return value, nil
}
