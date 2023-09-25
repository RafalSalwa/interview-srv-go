package tracing

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func RegisterMetricsEndpoint(addr string) {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		fmt.Println("Starting Prometheus metrics server")
		if err := http.ListenAndServe(addr, nil); err != nil {
			fmt.Println("Failed to serve Prometheus metrics:", err)
		}
	}()
}
