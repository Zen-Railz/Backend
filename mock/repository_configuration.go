package mock

import (
	"zenrailz/errorr"
)

func NewConfigurationRepository() *ConfigurationRepository {
	return &ConfigurationRepository{}
}

type ConfigurationRepository struct {
	list        map[string]string
	sourceError errorr.Entity
}

func (r *ConfigurationRepository) List(category string) (map[string]string, errorr.Entity) {
	return r.list, r.sourceError
}

func (r *ConfigurationRepository) Value(category string, key string) (string, errorr.Entity) {
	return r.list[key], r.sourceError
}

func (r *ConfigurationRepository) EmptyList() *ConfigurationRepository {
	r.list = make(map[string]string)
	return r
}

func (r *ConfigurationRepository) AddValue(key string, value string) *ConfigurationRepository {
	if r.list == nil {
		r.EmptyList()
	}

	r.list[key] = value
	return r
}

func (r *ConfigurationRepository) SetSourceError() *ConfigurationRepository {
	r.sourceError = NewError().
		SetCode(ErrorCode)
	return r
}
