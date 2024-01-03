package middleware

import (
	"net/http"
	"strconv"
	"time"

	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	requestLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "lybbrio",
			Name:      "http_request_duration_seconds",
			Help:      "Histogram of latencies for HTTP requests.",
			Buckets:   []float64{.25, .5, 1, 2.5, 5, 10},
		},
		[]string{"method", "path", "status"},
	)
	requests = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "lybbrio",
			Name:      "http_requests_total",
			Help:      "Total number of HTTP requests made.",
		},
		[]string{"method", "path", "status"},
	)
	request_errors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "lybbrio",
			Name:      "http_request_errors_total",
			Help:      "Total number of HTTP requests that resulted in an error.",
		},
		[]string{"method", "path", "status"},
	)
)

func Prometheus(reg prometheus.Registerer) func(next http.Handler) http.Handler {
	reg.MustRegister(
		requestLatency,
		requests,
		request_errors,
	)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()

			ww := chi_middleware.NewWrapResponseWriter(w, r.ProtoMajor) // Wrap response writer

			defer func() {
				requestLatency.WithLabelValues(
					r.Method, r.URL.Path, strconv.Itoa(ww.Status()),
				).Observe(time.Since(start).Seconds())

				requests.WithLabelValues(
					r.Method, r.URL.Path, strconv.Itoa(ww.Status()),
				).Inc()

				if ww.Status() >= 500 {
					request_errors.WithLabelValues(
						r.Method, r.URL.Path, strconv.Itoa(ww.Status()),
					).Inc()
				}
			}()

			next.ServeHTTP(ww, r)
		})
	}
}
