package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/v3/mem"
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
	Name: "minitwit_request_duration_milliseconds",
	Help: "Request duration distribution.",
})

var requestStart = time.Now()

func beforeRequestMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestStart = time.Now()
		context.Next()
	}
}

func afterRequestMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		RESPONSE_COUNTER.Inc()
		requestTime := time.Since(requestStart)
		REQUEST_DURATION_SUMMARY.Observe(float64(requestTime.Milliseconds()))

		v, _ := mem.VirtualMemory()
		CPU_LOAD.Set(v.UsedPercent)
	}
}

func ConfigureMetrics(router *gin.Engine) {
	router.Use(beforeRequestMiddleware())
	router.Use(afterRequestMiddleware())

	prometheus.MustRegister(RESPONSE_COUNTER)
	prometheus.MustRegister(REQUEST_DURATION_SUMMARY)
	prometheus.MustRegister(CPU_LOAD)
}

func MapMetricsEndpoints(router *gin.Engine) {
	router.GET("/metrics", metricsHandler())
}

func metricsHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
