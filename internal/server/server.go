package server

import (
	"fmt"
	"net/http"

	"github.com/loomchat/api-gateway-loom-chat/internal/config"
	"github.com/loomchat/api-gateway-loom-chat/internal/middleware"
	"github.com/loomchat/api-gateway-loom-chat/internal/routing"
	"github.com/loomchat/api-gateway-loom-chat/pkg/log"
)

var appConfigs = config.GetConfigs()
var serveMux = routing.SetUpServeMux()

var serverHandler = middleware.PrependMiddlewareChain(
	serveMux,
	middleware.RateLimitMiddleware,
	middleware.LogMiddleware,
	middleware.AuthMiddleware,
)

var server = &http.Server{
	Addr:           fmt.Sprintf(":%d", appConfigs.Port),
	Handler:        serverHandler,
	ReadTimeout:    appConfigs.Timeout,
	WriteTimeout:   appConfigs.Timeout,
	MaxHeaderBytes: 1 << 20,
}

func Start() error {
	env := config.GetEnv()
	log.Debug("Environment variables: %v", env.Variables)

	log.Debug("Raw app configs: %s", appConfigs)
	log.Debug("Replacing variable occurrences with their values in the app configs...")
	err := config.ReplaceEnvVarsInConfigs(appConfigs, env.Variables)
	if err != nil {
		log.Fatal("Failed to replace env variables occurrences with their values in the app configs")
		return err
	}
	log.Debug("Processed app configs: %s", appConfigs)

	log.Debug("Setting up the routes...")
	routing.SetUpRouteHandlers(appConfigs, env, serveMux)

	log.Info("Listening on port %d...", appConfigs.Port)
	return server.ListenAndServe()
}
