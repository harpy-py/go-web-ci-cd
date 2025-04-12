package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// RequestCount metric exported
var RequestCount *prometheus.CounterVec

func InitMetrics() {
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "go_app_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method"},
	)

	// Register
	prometheus.MustRegister(RequestCount)
}

// ExposeMetricsHandler returns an http.Handler for /metrics
func ExposeMetricsHandler() http.Handler {
	return promhttp.Handler()
}