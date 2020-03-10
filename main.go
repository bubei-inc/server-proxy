package main

import (
	"proxy/container/client"
	"proxy/container/server"
	"time"
)

func main() {

	go server.StartServer()

	go client.StartClient()

	for {
		time.Sleep(1 * time.Second)
	}
}
