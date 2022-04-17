package service

import (
	"zenrailz/errorr"
	"zenrailz/service/native"
	"zenrailz/service/railway"
)

type Native interface {
	SystemHealth() *native.SystemState
	DatabaseHealth() *native.DatabaseState
}

type Railway interface {
	Stations() (interface{}, errorr.Entity)
	Lines() ([]railway.Line, errorr.Entity)
}
