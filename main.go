package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	//Register prometheus
	Register()

	routes := mux.NewRouter()
	//Routes
	routes.HandleFunc("/", HomeHandler)
	routes.HandleFunc("/master", HomeHandler)
	routes.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", routes))
}
