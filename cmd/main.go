package main

import (
	"os"

	"github.com/loomchat/api-gateway-loom-chat/internal/server"
	"github.com/loomchat/api-gateway-loom-chat/pkg/log"
)

func main() {
	err := server.Start()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}
