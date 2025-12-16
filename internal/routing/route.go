package routing

import (
	"net/http"
)

func SetUpServeMux() *http.ServeMux {
	sv := http.NewServeMux()

	sv.HandleFunc("GET /hello", SendHello)

	return sv
}

func SendHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, friend!"))
}
