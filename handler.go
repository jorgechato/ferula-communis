package main

import (
	"fmt"
	"net/http"
	"time"
)

// HomeHandler is a Hello World view
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	fmt.Fprintln(w, "Hi there!, You are in ", r.URL.Path)

	duration := time.Since(start)
	go latencyEndpoint(r.URL.Path, http.StatusOK, duration)
}
