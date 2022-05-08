package configuration

import (
	"zenrailz/code"
	"zenrailz/errorr"
)

func (r *Repository) Value(category string, key string) (string, errorr.Entity) {
	configurations, err := r.List(category)
	if err != nil {
		return "", err.Trace()
	}

	value, valueExist := configurations[key]
	if valueExist {
		return value, nil
	} else {
		return "", errorr.New(code.ConfigRepoValueNotFound, "Unable to get value from configuration", map[string]string{
			"category": category,
			"key":      key,
		})
	}
}
