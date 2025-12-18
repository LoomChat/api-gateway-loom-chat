package routing

import (
	"fmt"
	"io"
	"net/http"

	"github.com/loomchat/api-gateway-loom-chat/internal/config"
	"github.com/loomchat/api-gateway-loom-chat/pkg/log"
)

func SetUpServeMux() *http.ServeMux {
	sv := http.NewServeMux()

	return sv
}

var httpClient = http.Client{} // TODO: SET UP THE CLIENT PROPERLY!!!

func SetUpRouteHandlers(
	appConfigs *config.Configs, env *config.Environment, serveMux *http.ServeMux,
) error {
	routes := config.GetRoutes(appConfigs)

	for _, route := range routes {
		log.Debug("Replacing variables with their values in a route: %s", route)
		err := config.ReplaceEnvVarsInConfigs(route, env.Variables)
		if err != nil {
			return err
		}
		log.Debug("Replaced route: %s", route)
	}

	for _, route := range routes {
		log.Debug("Adding a handler for route: %s", route)
		addRouteHandler(route, serveMux)
	}

	return nil
}

func getBackendUrl(r *config.Route) string {
	return fmt.Sprintf("%s://%s:%s%s", r.Proto, r.Host, r.Port, r.BackendEndpoint)
}

func addRouteHandler(route *config.Route, serveMux *http.ServeMux) {
	serveMux.HandleFunc(
		route.Endpoint,
		func(w http.ResponseWriter, r *http.Request) {
			log.Fixme("IMPLEMENT REQUEST PREPOCESSING!!!")
			req, err := http.NewRequest(route.Method, getBackendUrl(route), r.Body)
			if err != nil {
				log.Error("Failed to create a request for the http client for the route: %s", route)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			// Copy headers
			for key, values := range req.Header {
				for _, value := range values {
					req.Header.Add(key, value)
				}
			}

			resp, err := httpClient.Do(req)
			if err != nil {
				log.Error(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
			defer resp.Body.Close()

			log.Fixme("IMPLEMENT RESPONSE PREPROCESSING!!!")

			// Copy headers
			for key, values := range resp.Header {
				for _, value := range values {
					w.Header().Add(key, value)
				}
			}

			w.WriteHeader(resp.StatusCode)
			io.Copy(w, resp.Body)
		},
	)
}
