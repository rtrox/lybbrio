package config

import (
	"crypto/rand"
	"lybbrio/internal/ent/schema/argon2id"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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
	"argon2id": map[string]interface{}{
		"iterations":  3,
		"memory":      12288,
		"parallelism": 1,
		"keyLength":   32,
		"saltLength":  16,
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
	Driver          string        `koanf:"driver" validate:"oneof=sqlite3 mysql postgres"`
	DSN             string        `koanf:"dsn" validate:"required"`
	MaxIdleConns    int           `koanf:"max-idle-conns"`
	MaxOpenConns    int           `koanf:"max-open-conns"`
	ConnMaxLifetime time.Duration `koanf:"conn-max-lifetime"`
}

type TaskConfig struct {
	Workers     int           `koanf:"workers" validate:"number,gt=0"`
	QueueLength int           `koanf:"queue-length" validate:"number,gt=0"`
	Cadence     time.Duration `koanf:"cadence" validate:"required"`
}

type JWTConfig struct {
	SigningMethod string        `koanf:"signing-method" validate:"oneof=HS512 RS512"`
	Expiry        time.Duration `koanf:"expiry" validate:"required"`
	Issuer        string        `koanf:"issuer" validate:"url"`
	HMACSecret    string        `koanf:"hmac-secret"`
	RSAPrivateKey string        `koanf:"rsa-private-key" validate:"omitempty,file"`
	RSAPublicKey  string        `koanf:"rsa-public-key" validate:"omitempty,file"`
}

func ValidateJWTConfig(sl validator.StructLevel) {
	c := sl.Current().Interface().(JWTConfig)

	if c.SigningMethod == "HS512" {
		if c.HMACSecret == "" {
			sl.ReportError(c.HMACSecret, "hmac-secret", "HMACSecret", "required", "")
		}
	}
	if c.SigningMethod == "RS512" {
		if c.RSAPrivateKey == "" {
			sl.ReportError(c.RSAPrivateKey, "rsa-private-key", "RSAPrivateKey", "required", "")
		}
		if c.RSAPublicKey == "" {
			sl.ReportError(c.RSAPublicKey, "rsa-public-key", "RSAPublicKey", "required", "")
		}
	}
}

func ValidateArgon2IDConfig(sl validator.StructLevel) {
	c := sl.Current().Interface().(argon2id.Config)

	// https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#argon2id
	if c.Iterations == 3 && c.Memory < 12288 {
		sl.ReportError(c.Memory, "memory", "Memory", "min", "12288")
	}

	if c.Iterations == 4 && c.Memory < 9216 {
		sl.ReportError(c.Memory, "memory", "Memory", "min", "9216")
	}

	if c.Iterations == 5 && c.Memory < 7168 {
		sl.ReportError(c.Memory, "memory", "Memory", "min", "7168")
	}
}

type Config struct {
	// Logging
	LogLevel  string `koanf:"log-level" validate:"log_level"`
	LogFormat string `koanf:"log-format" validate:"oneof=json text"`
	// HTTPServer
	GracefulShutdownTimeout time.Duration `koanf:"graceful-shutdown-timeout"`
	BaseURL                 string        `koanf:"base-url" validate:"url"`
	Interface               string        `koanf:"interface" validate:"ip"`
	Port                    int           `koanf:"port" validate:"number,gt=0"`
	DevMode                 bool          `koanf:"dev-mode"`

	GoCollector      bool `koanf:"go-collector"`
	ProcessCollector bool `koanf:"process-collector"`

	CalibreLibraryPath string `koanf:"calibre-library-path" validate:"required"`

	DB       DatabaseConfig  `koanf:"db"`
	Task     TaskConfig      `koanf:"task"`
	JWT      JWTConfig       `koanf:"jwt"`
	Argon2ID argon2id.Config `koanf:"argon2id"`

	k *koanf.Koanf
}

func (c *Config) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterStructValidation(ValidateJWTConfig, JWTConfig{})
	validate.RegisterStructValidation(ValidateArgon2IDConfig, argon2id.Config{})
	err := validate.RegisterValidation("log_level", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		if _, err := zerolog.ParseLevel(value); err != nil {
			return false
		}
		return true
	})
	if err != nil {
		return err
	}
	return validate.Struct(c)
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
