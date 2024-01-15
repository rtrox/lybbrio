package config

import (
	"crypto/rand"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gookit/validate"
	flag "github.com/spf13/pflag"

	"github.com/knadh/koanf/v2"

	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/posflag"

	"github.com/rs/zerolog"
)

var defaultSettings = map[string]interface{}{
	"log-level":                 "info",
	"log-format":                "text",
	"graceful-shutdown-timeout": 1 * time.Second,
	"base-url":                  "http://localhost:8080",
	"port":                      8080,
	"interface":                 "127.0.0.1",
	"go-collector":              true,
	"process-collector":         true,
	"calibre-library-path":      "./books/",
	"db": map[string]interface{}{
		"driver":            "sqlite3",
		"dsn":               "file:database/app.db?cache=shared&_fk=1",
		"max-idle-conns":    10,
		"max-open-conns":    100,
		"conn-max-lifetime": 1 * time.Hour,
	},
	"task": map[string]interface{}{
		"workers":      10,
		"queue-length": 100, // 1000 tasks
		"cadence":      5 * time.Second,
	},
	"jwt": map[string]interface{}{
		"signing-method": "HS512",
		"issuer":         "http://localhost:8080",
		"expiry":         1 * time.Hour,
	},
}

func RegisterFlags(flagSet *flag.FlagSet) {
	// flagSet.String("config", "config.yaml", "Path to config file")
	flagSet.String("log-level", "info", "Log level")
	flagSet.String("log-format", "text", "Log format (text, json)")
	flagSet.Duration("graceful-shutdown-timeout", 1*time.Second, "Graceful shutdown timeout prior to killing the process")
	flagSet.String("base-url", "http://localhost:8080", "Base URL")
	flagSet.Bool("go-collector", false, "Enable Go prometheus collector")
	flagSet.Bool("process-collector", false, "Enable process prometheus collector")
	flagSet.Int("port", 8080, "Port to Listen On")
	flagSet.String("interface", "127.0.0.1", "Interface")
	flagSet.String("calibre-library-path", "./books/", "Path to Calibre database")
	flagSet.String("db.driver", "sqlite3", "Database driver")
	flagSet.String("db.dsn", "file:database/app.db?cache=shared&_fk=1", "Database DSN")
	flagSet.Int("db.max-idle-conns", 10, "Database max idle connections")
	flagSet.Int("db.max-open-conns", 100, "Database max open connections")
	flagSet.Duration("db.conn-max-lifetime", 1*time.Hour, "Database connection max lifetime")
	flagSet.Int("task.workers", 10, "Number of workers")
	flagSet.Int("task.queue-length", 100, "Task queue length")
	flagSet.Duration("task.cadence", 5*time.Second, "Task cadence")
	flagSet.String("jwt.issuer", "http://localhost:8080", "JWT Issuer")
	flagSet.Duration("jwt.expiry", 1*time.Hour, "JWT Expiry")
	flagSet.String("jwt.signing-method", "HS512", "JWT Signing Method")
	flagSet.String("jwt.hmac-secret", "", "JWT HMACSecret")
	flagSet.String("jwt.rsa-private-key", "", "JWT RSAPrivateKey")
	flagSet.String("jwt.rsa-public-key", "", "JWT RSAPublicKey")
}

type DatabaseConfig struct {
	Driver          string        `koanf:"driver" validate:"required|in:sqlite3,mysql,postgres"`
	DSN             string        `koanf:"dsn" validate:"required"`
	MaxIdleConns    int           `koanf:"max-idle-conns"`
	MaxOpenConns    int           `koanf:"max-open-conns"`
	ConnMaxLifetime time.Duration `koanf:"conn-max-lifetime"`
}

type TaskConfig struct {
	Workers     int           `koanf:"workers" validate:"required|int|gt:0"`
	QueueLength int           `koanf:"queue-length" validate:"required|int|gt:0"`
	Cadence     time.Duration `koanf:"cadence" validate:"required"`
}

type JWTConfig struct {
	SigningMethod string        `koanf:"signing-method" validate:"required|in:HS512,RS512"`
	Expiry        time.Duration `koanf:"expiry" validate:"required"`
	Issuer        string        `koanf:"issuer" validate:"required|url"`
	HMACSecret    string        `koanf:"hmac-secret"`
	RSAPrivateKey string        `koanf:"rsa-private-key"`
	RSAPublicKey  string        `koanf:"rsa-public-key"`
}

func (c JWTConfig) Validate() error {
	v := validate.Struct(c)
	if !v.Validate() {
		return v.Errors
	}
	if c.SigningMethod == "HS512" {
		if c.HMACSecret == "" {
			return fmt.Errorf("HMACSecret is required for HS512")
		}
	}
	if c.SigningMethod == "RS512" {
		if c.RSAPrivateKey == "" {
			return fmt.Errorf("RSAPrivateKey is required for RS512")
		}
		if c.RSAPublicKey == "" {
			return fmt.Errorf("RSAPublicKey is required for RS512")
		}
	}
	return nil
}

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

	CalibreLibraryPath string `koanf:"calibre-library-path" validate:"required"`

	DB   DatabaseConfig `koanf:"db"`
	Task TaskConfig     `koanf:"task"`
	JWT  JWTConfig      `koanf:"jwt"`

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
	vd := validate.Struct(c.DB)
	if !vd.Validate() {
		return vd.Errors
	}
	vt := validate.Struct(c.Task)
	if !vt.Validate() {
		return vt.Errors
	}
	vj := validate.Struct(c.JWT)
	if !vj.Validate() {
		return vj.Errors
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
	if err := k.Load(confmap.Provider(defaultSettings, "."), nil); err != nil {
		return nil, err
	}

	if k.Get("jwt.signing-method") == "HS512" && !k.Exists("jwt.hmac-secret") {
		u, err := uuid.NewRandomFromReader(rand.Reader)
		if err != nil {
			return nil, err
		}
		if err := k.Set("jwt.hmac-secret", u.String()); err != nil {
			return nil, err
		}
	}
	// Environment
	if err := k.Load(env.Provider("", ".", func(s string) string {
		return strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(s), "__", "."), "_", "-")
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
		"CalibreLibraryPath":      "calibre-library-path",
	}
}

func (c DatabaseConfig) Translates() map[string]string {
	return validate.MS{
		"Driver":          "driver",
		"DSN":             "dsn",
		"MaxIdleConns":    "max-idle-conns",
		"MaxOpenConns":    "max-open-conns",
		"ConnMaxLifetime": "conn-max-lifetime",
	}
}
