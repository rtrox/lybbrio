package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
)

func Test_PrometheusMiddleware(t *testing.T) {
	tests := []struct {
		name            string
		handler         http.Handler
		expectedMetrics []string
	}{
		{
			name: "200 OK",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("OK"))
				w.WriteHeader(http.StatusOK)
			}),
			expectedMetrics: []string{
				`lybbrio_http_request_duration_seconds`,
				`lybbrio_http_requests_total`,
			},
		},
		{
			name: "500 Internal Server Error",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			}),
			expectedMetrics: []string{
				`lybbrio_http_request_duration_seconds`,
				`lybbrio_http_requests_total`,
				`lybbrio_http_request_errors_total`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)

			reg := prometheus.NewPedanticRegistry()

			handler := Prometheus(reg)(tt.handler)

			req := httptest.NewRequest("GET", "/", nil)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			mfs, err := reg.Gather()
			require.Nil(err)
			require.Len(mfs, len(tt.expectedMetrics))
			for mf := range mfs {
				require.Contains(tt.expectedMetrics, *mfs[mf].Name)
			}
		})
	}
}
