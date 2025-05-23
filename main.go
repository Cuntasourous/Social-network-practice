package main

import (
	Server "social-network/backend/Server"
	database "social-network/backend/db/Database"
)

func main() {
	database.InitDB()
	Server.ServeDistFiles()
	Server.Routes()
	Server.Server()
}
