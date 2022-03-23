package environment

import "zenrailz/anomaly"

func ServerPort() (string, *anomaly.ServiceError) {
	return getValue("PORT")
}

func DatabaseUri() (string, *anomaly.ServiceError) {
	return getValue("DATABASE_URL")
}

func LogLevel() (string, *anomaly.ServiceError) {
	return getValue("LOG_LEVEL")
}
