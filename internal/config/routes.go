package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/loomchat/api-gateway-loom-chat/pkg/log"
)

// Routes struct represents the configurations parsed from the master file of the routes configurations.
type Routes struct {
	Files []string `json:"files"` // filenames of the files that contain routing configurations
}

// Route struct represents a route configuration information parsed from the routes configuration
// files.
// Fields:
//   - Host: a host of the underlying backend service that will be called by the api gateway.
//   - Port: a port of the underlying backend service that will be called by the api gateway.
//   - Endpoint: an endpoint of the api gateway.
//   - BackendEndpoint: the backend service endpoint that is mapped to the api gateway endpoint.
//   - Proto: the application protocol used for communication between the api gateway and the internal
//     backend servers.
//
// Any substring value that is preceded by '$', contains only one ore more alphanumeric and '_'
// (underscore) characters inside string values of the Route struct fields are considered to be an
// environment variable name.
type Route struct {
	Host            string `json:"host"`
	Port            string `json:"port"`
	Endpoint        string `json:"endpoint"`
	BackendEndpoint string `json:"backendEndpoint"`
	Proto           string `json:"proto"`
}

func (r Route) String() string {
	return fmt.Sprintf(
		"{host: %s, port: %s, endpoint: %s, backendEndpoint: %s, proto: %s",
		r.Host,
		r.Port,
		r.Endpoint,
		r.BackendEndpoint,
		r.Proto,
	)
}

func GetRoutes(appConfigs *Configs) []*Route {
	routes, err := parseRoutes(appConfigs)
	if err != nil {
		log.Fatal("Failed to parse routes configuration :(. No routes -> no api gateway...")
		log.Fatal("Exiting...")
		os.Exit(1)
	}

	return routes
}

func parseRoutes(appConfigs *Configs) ([]*Route, error) {
	routesMaster, err := parseRoutesMasterFile(appConfigs)
	if err != nil {
		return nil, err
	}

	var allRoutes []*Route
	for _, filename := range routesMaster.Files {
		routes, err := parseRoutesFile(filename)
		if err != nil {
			return nil, err
		}
		allRoutes = append(allRoutes, routes...)
	}

	return allRoutes, nil
}

func parseRoutesFile(filename string) ([]*Route, error) {
	routesBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Error("Failed to read from a routes file: %s", filename)
		return nil, err
	}

	var routes []*Route
	err = json.Unmarshal(routesBytes, &routes)
	if err != nil {
		log.Error("Failed to json unmarshal '%s' file content bytes due to: %s", filename, err)
		return nil, err
	}

	return routes, nil
}

func parseRoutesMasterFile(appConfigs *Configs) (*Routes, error) {
	routesMasterBytes, err := os.ReadFile(appConfigs.RouteConfigPath)
	if err != nil {
		log.Error(
			"Failed to read from the routes master file '%s' due to %s",
			appConfigs.RouteConfigPath,
			err,
		)
		return nil, err
	}

	var routesMaster Routes
	err = json.Unmarshal(routesMasterBytes, &routesMaster)
	if err != nil {
		log.Error("Failed to json unmarshal the routes master file content bytes due to: %s", err)
		return nil, err
	}

	return &routesMaster, nil
}
