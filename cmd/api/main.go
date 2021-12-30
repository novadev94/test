package main

import (
	"github.com/chikob3/test/server"
)

func main() {
	server := server.NewServer()
	server.Serve("8080")
}
