package metrics

import (
    "fmt"
    "github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	SuccessHttpRequests prometheus.Counter
	ErrorHttpRequests   prometheus.Counter
}

func NewApiGatewayMetrics(serviceName string) *Metrics {
	return &Metrics{
		SuccessHttpRequests: prometheus.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_http_requests_total", serviceName),
			Help: "The total number of success http requests",
		}),
		ErrorHttpRequests: prometheus.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_http_requests_total", serviceName),
			Help: "The total number of error http requests",
		}),
	}
}
