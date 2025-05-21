package main

import (
	Server "social-network/backend/Server"
)

func main() {
	Server.ServeDistFiles()
	Server.Routes()
	Server.Server()
}
