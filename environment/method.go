package environment

func ServerPort() (string, error) {
	return getValue("PORT")
}

func DatabaseUri() (string, error) {
	return getValue("DATABASE_URI")
}

func LogLevel() (string, error) {
	return getValue("LOG_LEVEL")
}
