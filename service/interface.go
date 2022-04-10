package service

import (
	"zenrailz/anomaly"
	"zenrailz/service/native"
)

type Native interface {
	SystemHealth() *native.SystemState
	DatabaseHealth() *native.DatabaseState
}

type Railway interface {
	Stations() (interface{}, *anomaly.ServiceError)
}
