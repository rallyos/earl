package main

import (
	"fmt"

	"github.com/shiftingphotons/earl/app"
)

func main() {
	server := new(app.Server)
	fmt.Println("Starting server...")
	server.Initialize()
}
