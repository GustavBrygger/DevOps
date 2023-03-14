package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var CPU_LOAD = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "minitwit_cpu_load_percent",
	Help: "Current load of the CPU in percent.",
})

var RESPONSE_COUNTER = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "minitwit_http_responses_total",
	Help: "The count of HTTP responses sent.",
})

var REQUEST_DURATION_SUMMARY = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:    "minitwit_request_duration_milliseconds",
	Help:    "Request duration distribution.",
	Buckets: prometheus.ExponentialBuckets(0.1, 1.5, 5),
})

var registry = prometheus.NewRegistry()

func ConfigurePrometheus() {
	registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	registry.MustRegister(RESPONSE_COUNTER)
	registry.MustRegister(REQUEST_DURATION_SUMMARY)
	registry.MustRegister(CPU_LOAD)
}

func MapMetricsEndpoints(router *gin.Engine) {
	router.GET("/metrics", metricsHandler())
}

func metricsHandler() gin.HandlerFunc {
	h := promhttp.HandlerFor(
		registry,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
			Registry:          registry,
		})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
