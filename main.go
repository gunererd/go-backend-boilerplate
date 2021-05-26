package main

import (
	"os"

	"gitlab.com/gunererd/dummy-challange/src"
)

func getPort() string {
	port := os.Getenv("PORT")

	if string(port[0]) != ":" {
		port = ":" + port
	}

	if port == "" {
		port = ":8080"
	}
	return port
}

func main() {

	port := getPort()

	server := src.NewServer()
	server.Run(port)
}
