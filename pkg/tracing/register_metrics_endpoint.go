package tracing

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
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
