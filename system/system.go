package system

func Start() {
	server := new(Server)
	server.Run()
}
