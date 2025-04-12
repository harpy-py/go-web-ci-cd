package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize metrics
	InitMetrics()

	// Routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RequestCount.WithLabelValues(r.Method).Inc()
		w.Write([]byte("Hello from modular Go app 👋"))
	})

	// Metrics route
	http.Handle("/metrics", ExposeMetricsHandler())

	log.Println("🚀 Go app running at :5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
