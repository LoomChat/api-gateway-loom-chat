package routing

import (
	"net/http"

	"github.com/loomchat/api-gateway-loom-chat/internal/config"
)

func SetUpServeMux() *http.ServeMux {
	sv := http.NewServeMux()

	sv.HandleFunc("GET /hello", SendHello)

	return sv
}

func setUpRoutes(appConfigs *config.Configs, serveMux *http.ServeMux) {
	// routes := config.GetRoutes(appConfigs)
	// env := config.GetEnv()
	//
	// for _, route := range routes {
	//
	// }
}

func SendHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, friend!"))
}
