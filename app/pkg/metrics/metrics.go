package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cast"
)

const (
	namespace = "goframe"
	subsystem = "ginBasic"
)

var metricsRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "requests_total",
		Help:      "request(ms) total",
	},
	[]string{"method", "path"},
)

var metricsRequestsCost = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "requests_cost",
		Help:      "request(ms) cost milliseconds",
	},
	[]string{"method", "path", "success", "http_code", "business_code", "runtime_milliseconds", "trace_id"},
)

func init() {
	prometheus.MustRegister(metricsRequestsTotal, metricsRequestsCost)
}

func RecordMetrics(method, uri string, success bool, httpCode, businessCode int, costSeconds float64, traceId string) {
	metricsRequestsTotal.With(prometheus.Labels{
		"method": method,
		"path":   uri,
	}).Inc()

	metricsRequestsCost.With(prometheus.Labels{
		"method":               method,
		"path":                 uri,
		"success":              cast.ToString(success),
		"http_code":            cast.ToString(httpCode),
		"business_code":        cast.ToString(businessCode),
		"runtime_milliseconds": cast.ToString(costSeconds * 1000),
		"trace_id":             traceId,
	}).Observe(costSeconds)
}
