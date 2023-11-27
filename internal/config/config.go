package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/gookit/validate"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/rs/zerolog"
	flag "github.com/spf13/pflag"
)

type Config struct {
	// Logging
	LogLevel  string `koanf:"log-level" validate:"ValidateLogLevel"`
	LogFormat string `koanf:"log-format" validate:"in:json,text"`
	// HTTPServer
	GracefulShutdownTimeout time.Duration `koanf:"graceful-shutdown-timeout"`
	BaseURL                 string        `koanf:"base-url" validate:"required|url"`
	Interface               string        `koanf:"interface" validate:"required|ip"`
	Port                    int           `koanf:"port" validate:"required|int|gt:0"`
	DevMode                 bool          `koanf:"dev-mode"`

	GoCollector      bool `koanf:"go-collector"`
	ProcessCollector bool `koanf:"process-collector"`

	k *koanf.Koanf
}

func (c Config) ValidateLogLevel(val string) bool {
	if _, err := zerolog.ParseLevel(val); err != nil {
		return false
	}
	return true
}

func (c *Config) Validate() error {
	v := validate.Struct(c)
	if !v.Validate() {
		return v.Errors
	}
	return nil
}

func (c Config) Messages() map[string]string {
	lvls := []string{}
	for i := zerolog.TraceLevel; i < zerolog.Disabled; i++ {
		if i.String() != "" {
			lvls = append(lvls, i.String())
		}
	}
	return validate.MS{
		"LogLevel.ValidateLogLevel": fmt.Sprintf("Invalid log level, must be one of: [%s]", strings.Join(lvls, ", ")),
	}
}

func LoadConfig(flagSet *flag.FlagSet) (*Config, error) {
	k := koanf.New(".")

	// Defaults
	if err := k.Load(confmap.Provider(map[string]interface{}{
		"log-level":                 "info",
		"log-format":                "text",
		"graceful-shutdown-timeout": "1s",
		"base-url":                  "http://localhost:8080",
		"port":                      8080,
		"interface":                 "127.0.0.1",
	}, "."), nil); err != nil {
		return nil, err
	}

	// Environment
	if err := k.Load(env.Provider("", ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(s), "_", "-")
	}), nil); err != nil {
		return nil, err
	}

	if err := k.Load(posflag.Provider(flagSet, ".", k), nil); err != nil {
		return nil, err
	}

	var out *Config
	if err := k.Unmarshal("", &out); err != nil {
		return nil, err
	}

	out.k = k
	return out, nil
}

func (c Config) Translates() map[string]string {
	return validate.MS{
		"LogLevel":                "log-level",
		"LogFormat":               "log-format",
		"GracefulShutdownTimeout": "graceful-shutdown-timeout",
		"BaseURL":                 "base-url",
		"GoCollector":             "go-collector",
		"ProcessCollector":        "process-collector",
		"Port":                    "port",
		"Interface":               "interface",
	}
}
