package metrics

import (
	"io"
	"regexp"
	"time"

	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/require"

	"github.com/tj/assert"
)

func TestAllMetricsPopulated(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	info := AppInfoGaugeFunc(AppInfoOpts{
		Namespace: "speedtest_exporter",
		Name:      "asdf",
		Version:   "v1.1.1",
		BuildTime: "2021-01-01T00:00:00Z",
		Revision:  "acb12356bd",
	})
	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(info)
	srv := httptest.NewServer(promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	defer srv.Close()

	c := &http.Client{Timeout: 300 * time.Millisecond}
	resp, err := c.Get(srv.URL)
	require.Nil(err)
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	require.Nil(err)

	/*
		# HELP speedtest_exporter_info Info about this application
		# TYPE speedtest_exporter_info gauge
		speedtest_exporter_info{app_name="<appName>",app_version="<appVersion>"} 1
	*/
	tests := []struct {
		desc  string
		match *regexp.Regexp
	}{
		{"device_info_desc", regexp.MustCompile(`(?m)^# HELP speedtest_exporter_info .*[a-zA-Z]+.*$`)},
		{"device_info", regexp.MustCompile(`(?m)^speedtest_exporter_info{app_name="asdf",app_version="v1.1.1",build_time="2021-01-01T00:00:00Z",revision="acb12356bd"} 1$`)},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert.True(tt.match.Match(buf), "Regex %s didn't match a line! Buffer: %s", tt.match.String(), buf)
		})
	}
}
