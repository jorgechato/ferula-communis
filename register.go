package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	//Track the time of any endpoint response
	endpointLatencies = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "adidas",
			Subsystem: "http_server",
			Name:      "http_responses_total",
			Help:      "The count of http responses issued, classified by code and method.",
		},
		[]string{
			"endpoint",
			"code",
		},
	)
)

//Register all the Prometheus variables
func Register() {
	prometheus.Register(endpointLatencies)
}
