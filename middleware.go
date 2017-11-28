package main

import (
	//"time"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if path := r.URL.Path; path != "/metrics" {
			go countEndpoint(path)
		}

		next.ServeHTTP(w, r)
	})

}

func countEndpoint(path string) {
	endpointCount.With(
		prometheus.Labels{
			"endpoint": path,
		},
	).Inc()
}
