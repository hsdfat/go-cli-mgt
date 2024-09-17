package main

import "go-cli-mgt/pkg/server"

func main() {
	server := server.NewFiber()

	server.Listen("localhost:8080")
}
