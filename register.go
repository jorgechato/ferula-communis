package main

import "github.com/prometheus/client_golang/prometheus"

var (
	endpointCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "adidas",
			Name:      "http_endpoint_counter",
			Help:      "Number of request on each endpoint",
		},
		[]string{
			"endpoint",
		},
	)
)

func Register() {
	prometheus.Register(endpointCount)
}
