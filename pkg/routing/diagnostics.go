package routing

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func DiagnosticsRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/healthz", healthzHandler()).Methods(http.MethodGet)
	r.HandleFunc("/readyz", readyzHandler()).Methods(http.MethodGet)
	return r
}

func healthzHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request is processing: %s", r.URL.Path)
		fmt.Fprint(w, "Health ok")
	}
}

func readyzHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request is processing: %s", r.URL.Path)
		fmt.Fprint(w, "Totally ready")
	}
}
