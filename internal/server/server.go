package server

import (
	"fmt"
	"net/http"

	"github.com/loomchat/api-gateway-loom-chat/internal/config"
	"github.com/loomchat/api-gateway-loom-chat/internal/middleware"
	"github.com/loomchat/api-gateway-loom-chat/internal/routing"
	"github.com/loomchat/api-gateway-loom-chat/pkg/log"
)

var serverConfigs = config.GetConfigs()

var serverHandler = middleware.PrependMiddlewareChain(
	routing.SetUpServeMux(),
	middleware.RateLimitMiddleware,
	middleware.LogMiddleware,
	middleware.AuthMiddleware,
)

var server = &http.Server{
	Addr:           fmt.Sprintf(":%d", serverConfigs.Port),
	Handler:        serverHandler,
	ReadTimeout:    serverConfigs.Timeout,
	WriteTimeout:   serverConfigs.Timeout,
	MaxHeaderBytes: 1 << 20,
}

func Start() {
	log.Info("Listening on port %d...", serverConfigs.Port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("%s", err)
	}
}
