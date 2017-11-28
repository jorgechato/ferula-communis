package main

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func latencyEndpoint(path string, code int, duration time.Duration) {
	endpointLatencies.With(
		prometheus.Labels{
			"endpoint": path,
			"code":     strconv.Itoa(code),
		},
	).Observe(duration.Seconds())
}
