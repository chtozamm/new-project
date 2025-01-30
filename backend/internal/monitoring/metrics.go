package monitoring

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var httpRequestCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "http_request_total",
	Help: "The total number of received HTTP requests",
}, []string{"status", "path", "method"})

var activeConnectionsGauge = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "active_connections",
	Help: "Number of active connections to the service",
})

var latencyHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "http_request_duration_seconds",
	Help:    "Duration of HTTP requests",
	Buckets: []float64{0.1, 0.5, 1, 2.5, 5, 10},
}, []string{"status", "path", "method"})

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (sr *statusRecorder) WriteHeader(code int) {
	sr.statusCode = code
	sr.ResponseWriter.WriteHeader(code)
}

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		activeConnectionsGauge.Inc()
		defer activeConnectionsGauge.Dec()

		now := time.Now()

		recorder := &statusRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(recorder, r)

		status := strconv.Itoa(recorder.statusCode)
		path := r.URL.Path
		method := r.Method

		latencyHistogram.With(prometheus.Labels{"method": method, "path": path, "status": status}).Observe(time.Since(now).Seconds())
		httpRequestCounter.WithLabelValues(status, path, method).Inc()
	})
}

func PrometheusHandler() http.Handler {
	reg := prometheus.NewRegistry()
	reg.MustRegister(httpRequestCounter)
	reg.MustRegister(activeConnectionsGauge)
	reg.MustRegister(latencyHistogram)
	return promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
}
