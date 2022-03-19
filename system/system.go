package system

import "zenrailz/log"

func Start() {
	defer recuperate()

	log.SetLevel()

	server := new(Server)
	server.Run()
}

func recuperate() {
	if err := recover(); err != nil {
		log.Error("Panic", err)
	}
}
