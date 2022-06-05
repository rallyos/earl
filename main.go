package main

import (
	"earl/app"
	"fmt"
)

func main() {
	server := new(app.Server)
	fmt.Println("Starting server...")
	server.Initialize()
}
