package server

import (
	"net/http"
	"time"
	"fmt"

	"github.com/loomchat/api-gateway-loom-chat/pkg/log"
)

const (
	Port = 8080 // TODO: GET IT FROM THE CONFIGS!!!
)

var server = &http.Server{
	Addr: fmt.Sprintf(":%d", Port),
	Handler: nil, // TODO: IMPLEMENT THE HANDLER!!!
	ReadTimeout: 10 * time.Second, // TODO: SHOULD BE TAKEN FROM THE CONFIGS AS WELL
	WriteTimeout: 10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}

func Start() {
	log.Info("Listening on port %d...", Port)
	server.ListenAndServe()
}
