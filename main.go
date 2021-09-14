package main

import (
	"estudar.com.br/config/database"
	"estudar.com.br/config/server"
)

// @title Golang
// @version 1.0
// @description Api de estudo Golang.

// @contact.name Glauber
// @contact.email glauber.barboza@hotmail.com

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	database.StartDB()
	server := server.NewServer()
	server.Run()
}
