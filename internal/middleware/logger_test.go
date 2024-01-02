package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

type logHook struct {
	Count int
}

func (logHook *logHook) Run(logEvent *zerolog.Event, level zerolog.Level, message string) {
	logHook.Count++
}

func Test_StructuredLogger(t *testing.T) {
	tests := []struct {
		name         string
		handler      http.Handler
		expectedFunc func() map[string]interface{}
	}{
		{
			name: "200 OK",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("OK"))
				w.WriteHeader(http.StatusOK)
			}),
			expectedFunc: func() map[string]interface{} {
				ret := make(map[string]interface{})
				ret["client_id"] = "192.0.2.1:1234" // httptest default client id - TEST-NET" in RFC 5737
				ret["method"] = "GET"
				ret["status_code"] = float64(200)
				ret["bytes_written"] = float64(2)
				ret["path"] = "/"
				ret["level"] = "info"
				return ret
			},
		},
		{
			name: "500 Internal Server Error",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			}),
			expectedFunc: func() map[string]interface{} {
				ret := make(map[string]interface{})
				ret["client_id"] = "192.0.2.1:1234" // httptest default client id - TEST-NET" in RFC 5737
				ret["method"] = "GET"
				ret["status_code"] = float64(500)
				ret["bytes_written"] = float64(0)
				ret["path"] = "/"
				ret["level"] = "error"
				return ret
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)
			out := &bytes.Buffer{}
			logger := zerolog.New(out)
			handle := StructuredLogger(&logger)(tt.handler)
			handle.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil))

			logLine := make(map[string]interface{})
			require.NoError(
				json.Unmarshal([]byte(out.String()), &logLine),
			)
			require.NotEqual("", logLine["latency"])
			for k, v := range tt.expectedFunc() {
				require.Equal(v, logLine[k])
			}
		})
	}
}

func Test_StructuredLoggerWithExcludedPath(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		shouldLog int
	}{
		{
			name:      "exclude",
			path:      "/health",
			shouldLog: 0,
		},
		{
			name:      "dont exclude",
			path:      "/",
			shouldLog: 1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()
			require := require.New(t)
			out := &bytes.Buffer{}
			logger := zerolog.New(out)
			hook := &logHook{}
			logger = logger.Hook(hook)
			handle := StructuredLogger(&logger, "/health")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))
			handle.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET", tt.path, nil))
			require.Equal(tt.shouldLog, hook.Count)
		})
	}
}

func Test_StructeredLoggerPathLogging(t *testing.T) {
	tests := []struct {
		name string
		path string
	}{
		{
			name: "root",
			path: "/",
		},
		{
			name: "subpath",
			path: "/subpath",
		},
		{
			name: "subpath with query",
			path: "/subpath?query=1",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			out := &bytes.Buffer{}
			logger := zerolog.New(out)
			handle := StructuredLogger(&logger)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))

			handle.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET", tt.path, nil))
			logLine := make(map[string]interface{})
			require.NoError(
				json.Unmarshal([]byte(out.String()), &logLine),
			)
			require.Equal(tt.path, logLine["path"])
		})
	}
}
