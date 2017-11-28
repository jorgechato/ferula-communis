package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// HomeHandler is a Hello World view
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	defer func() {
		duration := time.Since(start)
		go latencyEndpoint(r.URL.Path, http.StatusOK, duration)
	}()

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Fprintln(w, "Hi there!, You are in ", r.URL.Path)
}
