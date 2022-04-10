package service

import (
	"zenrailz/errorr"
	"zenrailz/service/native"
)

type Native interface {
	SystemHealth() *native.SystemState
	DatabaseHealth() *native.DatabaseState
}

type Railway interface {
	Stations() (interface{}, errorr.Entity)
}
