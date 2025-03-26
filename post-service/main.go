package main

import (
	"fmt"
	"post-service/config"
	"post-service/db"
	"post-service/server"
)

func main() {
	fmt.Println("Starting Post Service...")

	config.LoadEnv()
	db.Connect()

	server.StartServer()
}
