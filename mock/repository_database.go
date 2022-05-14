package mock

import "zenrailz/errorr"

func NewDatabaseRepository() *DatabaseRepository {
	return &DatabaseRepository{}
}

func (d *DatabaseRepository) Ping() errorr.Entity {
	return d.pingError
}

func (d *DatabaseRepository) SetPingError() *DatabaseRepository {
	d.pingError = NewError().
		SetCode(ErrorCode).
		SetStackTraceMessage(ErrorStackTraceMessage)
	return d
}

type DatabaseRepository struct {
	pingError errorr.Entity
}
