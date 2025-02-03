package logger

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var registerMetricsOnce sync.Once
var errorCounter *prometheus.CounterVec

func registerMetrics(namespace string) {
	registerMetricsOnce.Do(func() {
		errorCounter = promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "log_errors_total",
				Help:      "Total count of logged errors",
			},
			[]string{},
		)
	})
}
