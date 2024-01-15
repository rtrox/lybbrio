package config

import (
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	flag "github.com/spf13/pflag"
	"github.com/stretchr/testify/require"
)

func testFlagSet() *flag.FlagSet {
	flagSet := flag.NewFlagSet("test", flag.ContinueOnError)
	RegisterFlags(flagSet)
	return flagSet
}

func Test_LoadConfig_Defaults(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	config, err := LoadConfig(&flag.FlagSet{})
	require.NoError(err)

	require.Equal(defaultSettings["log-level"], config.LogLevel)
	require.Equal(defaultSettings["log-format"], config.LogFormat)
	require.Equal(defaultSettings["graceful-shutdown-timeout"], config.GracefulShutdownTimeout)
	require.Equal(defaultSettings["base-url"], config.BaseURL)
	require.Equal(defaultSettings["port"], config.Port)
	require.Equal(defaultSettings["interface"], config.Interface)
	require.Equal(defaultSettings["go-collector"], config.GoCollector)
	require.Equal(defaultSettings["process-collector"], config.ProcessCollector)
	require.Equal(defaultSettings["calibre-library-path"], config.CalibreLibraryPath)

	db := defaultSettings["db"].(map[string]interface{})
	require.Equal(db["driver"], config.DB.Driver)
	require.Equal(db["dsn"], config.DB.DSN)
	require.Equal(db["max-idle-conns"], config.DB.MaxIdleConns)
	require.Equal(db["max-open-conns"], config.DB.MaxOpenConns)
	require.Equal(db["conn-max-lifetime"], config.DB.ConnMaxLifetime)

	task := defaultSettings["task"].(map[string]interface{})
	require.Equal(task["workers"], config.Task.Workers)
	require.Equal(task["queue-length"], config.Task.QueueLength)
	require.Equal(task["cadence"], config.Task.Cadence)

	defaultJWT := defaultSettings["jwt"].(map[string]interface{})
	require.Equal(defaultJWT["issuer"], config.JWT.Issuer)
	require.Equal(defaultJWT["expiry"], config.JWT.Expiry)

	require.NotEmpty(config.JWT.HMACSecret, "JWT Secret should be set by default")
	firstSecret := config.JWT.HMACSecret

	config2, err := LoadConfig(&flag.FlagSet{})
	require.NoError(err)
	require.NotEqual(firstSecret, config2.JWT.HMACSecret, "JWT Secret should be unique when not set")
}

// nolint: errcheck
func Test_LoadConfig_Flags(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	flags := testFlagSet()
	flags.Set("log-level", "debug")
	flags.Set("log-format", "json")
	flags.Set("graceful-shutdown-timeout", "3s")
	flags.Set("base-url", "http://localhost:7070")
	flags.Set("go-collector", "false")
	flags.Set("process-collector", "false")
	flags.Set("port", "7070")
	flags.Set("interface", "192.0.2.1") // "TEST-NET" in RFC 5737
	flags.Set("calibre-library-path", "/tmp/")
	flags.Set("db.driver", "postgres")
	flags.Set("db.dsn", "postgres://user:pass@localhost:5432/dbname?sslmode=disable")
	flags.Set("db.max-idle-conns", "20")
	flags.Set("db.max-open-conns", "21")
	flags.Set("db.conn-max-lifetime", "2h")
	flags.Set("task.workers", "15")
	flags.Set("task.queue-length", "10")
	flags.Set("task.cadence", "1s")
	flags.Set("jwt.hmac-secret", "asdf")
	flags.Set("jwt.issuer", "issuer")
	flags.Set("jwt.expiry", "12h")

	config, err := LoadConfig(flags)
	require.NoError(err)

	require.Equal("debug", config.LogLevel)
	require.Equal("json", config.LogFormat)
	require.Equal(3*time.Second, config.GracefulShutdownTimeout)
	require.Equal("http://localhost:7070", config.BaseURL)
	require.False(config.GoCollector)
	require.False(config.ProcessCollector)
	require.Equal(7070, config.Port)
	require.Equal("/tmp/", config.CalibreLibraryPath)
	require.Equal("postgres", config.DB.Driver)
	require.Equal("postgres://user:pass@localhost:5432/dbname?sslmode=disable", config.DB.DSN)
	require.Equal(20, config.DB.MaxIdleConns)
	require.Equal(21, config.DB.MaxOpenConns)
	require.Equal(2*time.Hour, config.DB.ConnMaxLifetime)
	require.Equal(15, config.Task.Workers)
	require.Equal(10, config.Task.QueueLength)
	require.Equal(1*time.Second, config.Task.Cadence)
	require.Equal("asdf", config.JWT.HMACSecret)
	require.Equal("issuer", config.JWT.Issuer)
	require.Equal(12*time.Hour, config.JWT.Expiry)
}

func Test_LoadConfig_Env(t *testing.T) {
	// Env cannot be set in parallel tests
	require := require.New(t)

	t.Setenv("LOG_LEVEL", "debug")
	t.Setenv("LOG_FORMAT", "json")
	t.Setenv("GRACEFUL_SHUTDOWN_TIMEOUT", "3s")
	t.Setenv("BASE_URL", "http://localhost:7070")
	t.Setenv("GO_COLLECTOR", "false")
	t.Setenv("PROCESS_COLLECTOR", "false")
	t.Setenv("PORT", "7070")
	t.Setenv("INTERFACE", "192.0.2.1") // "TEST-NET" in RFC 5737
	t.Setenv("CALIBRE_LIBRARY_PATH", "/tmp/")
	t.Setenv("DB__DRIVER", "postgres")
	t.Setenv("DB__DSN", "postgres://user:pass@localhost:5432/dbname?sslmode=disable")
	t.Setenv("DB__MAX_IDLE_CONNS", "20")
	t.Setenv("DB__MAX_OPEN_CONNS", "21")
	t.Setenv("DB__CONN_MAX_LIFETIME", "2h")
	t.Setenv("TASK__WORKERS", "15")
	t.Setenv("TASK__QUEUE_LENGTH", "10")
	t.Setenv("TASK__CADENCE", "1s")
	t.Setenv("JWT__HMAC_SECRET", "asdf")
	t.Setenv("JWT__ISSUER", "issuer")
	t.Setenv("JWT__EXPIRY", "12h")

	config, err := LoadConfig(&flag.FlagSet{})
	require.NoError(err)

	require.Equal("debug", config.LogLevel)
	require.Equal("json", config.LogFormat)
	require.Equal(3*time.Second, config.GracefulShutdownTimeout)
	require.Equal("http://localhost:7070", config.BaseURL)
	require.False(config.GoCollector)
	require.False(config.ProcessCollector)
	require.Equal(7070, config.Port)
	require.Equal("/tmp/", config.CalibreLibraryPath)
	require.Equal("postgres", config.DB.Driver)
	require.Equal("postgres://user:pass@localhost:5432/dbname?sslmode=disable", config.DB.DSN)
	require.Equal(20, config.DB.MaxIdleConns)
	require.Equal(21, config.DB.MaxOpenConns)
	require.Equal(2*time.Hour, config.DB.ConnMaxLifetime)
	require.Equal(15, config.Task.Workers)
	require.Equal(10, config.Task.QueueLength)
	require.Equal(1*time.Second, config.Task.Cadence)
	require.Equal("asdf", config.JWT.HMACSecret)
	require.Equal("issuer", config.JWT.Issuer)
	require.Equal(12*time.Hour, config.JWT.Expiry)
}

// nolint: errcheck
func Test_LoadConfig_Env_Partials(t *testing.T) {
	// env cannot be set in parallel tests
	require := require.New(t)
	flags := testFlagSet()
	flags.Set("db.driver", "postgres")
	flags.Set("db.dsn", "postgres://user:pass@localhost:5432/dbname?sslmode=disable")
	flags.Set("db.max-idle-conns", "25")
	flags.Set("log-level", "debug")

	t.Setenv("DB__MAX_IDLE_CONNS", "20")
	t.Setenv("LOG_FORMAT", "json")

	config, err := LoadConfig(flags)
	require.NoError(err)

	require.Equal("postgres", config.DB.Driver, "flag should override default")
	require.Equal("postgres://user:pass@localhost:5432/dbname?sslmode=disable", config.DB.DSN, "flag should override defatul")
	require.Equal(25, config.DB.MaxIdleConns, "flag should override env and default")
	require.Equal("debug", config.LogLevel, "flag should override default")
	require.Equal("json", config.LogFormat, "env shoudl override default")
	require.Equal(10, config.Task.Workers, "should use default value")
}

func Test_Validate(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{
			name: "valid",
			config: Config{
				LogLevel:                "debug",
				LogFormat:               "json",
				GracefulShutdownTimeout: 3 * time.Second,
				BaseURL:                 "http://localhost:7070",
				Port:                    7070,
				Interface:               "192.0.2.1", // "TEST-NET" in RFC 5737
				GoCollector:             false,
				ProcessCollector:        false,
				CalibreLibraryPath:      "/tmp/",
				DB:                      DatabaseConfig{Driver: "sqlite3", DSN: "file::memory:?cache=shared"},
				Task:                    TaskConfig{Workers: 10, QueueLength: 100, Cadence: 5 * time.Second},
				JWT: JWTConfig{
					SigningMethod: "HS512",
					HMACSecret:    "asdf",
					Issuer:        "http://issuer.com",
					Expiry:        12 * time.Hour,
				},
			},
			wantErr: false,
		},
		{
			name: "invalid log level",
			config: Config{
				LogLevel:                "invalid",
				LogFormat:               "json",
				GracefulShutdownTimeout: 3 * time.Second,
				BaseURL:                 "http://localhost:7070",
				Port:                    7070,
				Interface:               "192.0.2.1", // "TEST-NET" in RFC 5737
				GoCollector:             false,
				ProcessCollector:        false,
				CalibreLibraryPath:      "/tmp/",
				DB:                      DatabaseConfig{Driver: "sqlite3", DSN: "file::memory:?cache=shared"},
				Task:                    TaskConfig{Workers: 10, QueueLength: 100, Cadence: 5 * time.Second},
				JWT: JWTConfig{
					SigningMethod: "HS512",
					HMACSecret:    "asdf",
					Issuer:        "http://issuer.com",
					Expiry:        12 * time.Hour,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid log format",
			config: Config{
				LogLevel:                "debug",
				LogFormat:               "invalid",
				GracefulShutdownTimeout: 3 * time.Second,
				BaseURL:                 "http://localhost:7070",
				Port:                    7070,
				Interface:               "192.0.2.1", // "TEST-NET" in RFC 5737
				GoCollector:             false,
				ProcessCollector:        false,
				CalibreLibraryPath:      "/tmp/",
				DB:                      DatabaseConfig{Driver: "sqlite3", DSN: "file::memory:?cache=shared"},
				Task:                    TaskConfig{Workers: 10, QueueLength: 100, Cadence: 5 * time.Second},
				JWT: JWTConfig{
					SigningMethod: "HS512",
					HMACSecret:    "asdf",
					Issuer:        "http://issuer.com",
					Expiry:        12 * time.Hour,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid db driver",
			config: Config{
				LogLevel:                "debug",
				LogFormat:               "text",
				GracefulShutdownTimeout: 3 * time.Second,
				BaseURL:                 "http://localhost:7070",
				Port:                    7070,
				Interface:               "192.0.2.1", // "TEST-NET" in RFC 5737
				GoCollector:             false,
				ProcessCollector:        false,
				CalibreLibraryPath:      "/tmp/",
				DB:                      DatabaseConfig{Driver: "invalid", DSN: "file::memory:?cache=shared"},
				Task:                    TaskConfig{Workers: 10, QueueLength: 100, Cadence: 5 * time.Second},
				JWT: JWTConfig{
					SigningMethod: "HS512",
					HMACSecret:    "asdf",
					Issuer:        "http://issuer.com",
					Expiry:        12 * time.Hour,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid db dsn",
			config: Config{
				LogLevel:                "debug",
				LogFormat:               "text",
				GracefulShutdownTimeout: 3 * time.Second,
				BaseURL:                 "http://localhost:7070",
				Port:                    7070,
				Interface:               "192.0.2.1", // "TEST-NET" in RFC 5737
				GoCollector:             false,
				ProcessCollector:        false,
				CalibreLibraryPath:      "/tmp/",
				DB:                      DatabaseConfig{Driver: "invalid", DSN: ""},
				Task:                    TaskConfig{Workers: 10, QueueLength: 100, Cadence: 5 * time.Second},
				JWT: JWTConfig{
					SigningMethod: "HS512",
					HMACSecret:    "asdf",
					Issuer:        "http://issuer.com",
					Expiry:        12 * time.Hour,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid task workers",
			config: Config{
				LogLevel:                "debug",
				LogFormat:               "text",
				GracefulShutdownTimeout: 3 * time.Second,
				BaseURL:                 "http://localhost:7070",
				Port:                    7070,
				Interface:               "192.0.2.1", // "TEST-NET" in RFC 5737
				GoCollector:             false,
				ProcessCollector:        false,
				CalibreLibraryPath:      "/tmp/",
				DB:                      DatabaseConfig{Driver: "invalid", DSN: ""},
				Task:                    TaskConfig{Workers: -1, QueueLength: 100, Cadence: 5 * time.Second},
				JWT: JWTConfig{
					SigningMethod: "HS512",
					HMACSecret:    "asdf",
					Issuer:        "http://issuer.com",
					Expiry:        12 * time.Hour,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)
			err := tt.config.Validate()
			if tt.wantErr {
				require.Error(err)
			} else {
				require.NoError(err)
			}
		})
	}
}

func TestJWTValidate(t *testing.T) {
	tests := []struct {
		name    string
		config  JWTConfig
		wantErr bool
	}{
		{
			name: "valid hs512",
			config: JWTConfig{
				SigningMethod: "HS512",
				Expiry:        12 * time.Hour,
				Issuer:        "http://issuer.com",
				HMACSecret:    "asdf",
			},
			wantErr: false,
		},
		{
			name: "valid rs512",
			config: JWTConfig{
				SigningMethod: "RS512",
				Expiry:        12 * time.Hour,
				Issuer:        "http://issuer.com",
				RSAPublicKey:  "./config.go", // needs to be existing file
				RSAPrivateKey: "./config.go", // needs to be existing file
			},
			wantErr: false,
		},
		{
			name: "invalid signing method",
			config: JWTConfig{
				SigningMethod: "HSDF",
				Expiry:        12 * time.Hour,
				Issuer:        "http://issuer.com",
				HMACSecret:    "asdf",
			},
			wantErr: true,
		},
		{
			name: "hs512 needs secret",
			config: JWTConfig{
				SigningMethod: "HS512",
				Expiry:        12 * time.Hour,
				Issuer:        "http://issuer.com",
			},
			wantErr: true,
		},
		{
			name: "rs512 needs signing key",
			config: JWTConfig{
				SigningMethod: "RS512",
				Expiry:        12 * time.Hour,
				Issuer:        "http://issuer.com",
				RSAPublicKey:  "asdf",
			},
			wantErr: true,
		},
		{
			name: "rs512 needs verifying key",
			config: JWTConfig{
				SigningMethod: "RS512",
				Expiry:        12 * time.Hour,
				Issuer:        "http://issuer.com",
				RSAPrivateKey: "asdf",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)
			validate := validator.New(validator.WithRequiredStructEnabled())
			validate.RegisterStructValidation(ValidateJWTConfig, JWTConfig{})
			err := validate.Struct(tt.config)
			if tt.wantErr {
				require.Error(err)
			} else {
				require.NoError(err)
			}
		})
	}
}
