package metrics

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var ()

type AppInfoOpts struct {
	Namespace string
	Name      string
	Version   string
	BuildTime string
	Revision  string
}

func AppInfoGaugeFunc(opts AppInfoOpts) prometheus.GaugeFunc {
	if opts.Namespace == "" {
		opts.Namespace = strings.ReplaceAll(opts.Name, "-", "_")
	}

	infoMetricOpts := prometheus.GaugeOpts{
		Namespace: opts.Namespace,
		Name:      "info",
		Help:      "Info about this application",
		ConstLabels: prometheus.Labels{
			"app_name":    opts.Name,
			"app_version": opts.Version,
			"build_time":  opts.BuildTime,
			"revision":    opts.Revision,
		},
	}
	return prometheus.NewGaugeFunc(
		infoMetricOpts,
		func() float64 { return 1 },
	)
}
