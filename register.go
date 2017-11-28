package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Track the time of any endpoint response
	endpointLatencies = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "adidas",
			Name:      "http_responses_Latency",
			Help:      "The latency, classified by code and endpoint.",
		},
		[]string{
			"endpoint",
			"code",
		},
	)
)

// Register all the Prometheus variables
func Register() {
	prometheus.Register(endpointLatencies)
}
