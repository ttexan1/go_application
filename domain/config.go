package domain

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config is the application config
type Config struct {
	Driver        string
	Env           string
	JWTSigningKey string
	Port          string
	PostgresURL   string
	SessionSecret string
}

// NewConfig returns the config
func NewConfig(envPath string) (*Config, error) {
	e, err := newEnv(envPath)
	if err != nil {
		return nil, err
	}
	return &Config{
		Driver:        e.get("DB_DRIVER", "postgres"),
		Env:           e.get("ENV", "development"),
		JWTSigningKey: e.get("JWT_SIGNING_KEY", "Development"),
		Port:          e.get("PORT", ":9000"),
		PostgresURL:   e.get("POSTGRES_URL", "host=localhost dbname=golang_practice_development sslmode=disable"),
		SessionSecret: e.get("SESSION_SECRET", "SUPER_SECRET"),
	}, nil
}

// BaseURL returns the base url for the env
func (c *Config) BaseURL() string {
	if c.IsProd() {
		return "http://golang-practice"
	}
	return "http://localhost" + c.Port
}

// IsDevelopment returns whether the env is development or not
func (c *Config) IsDevelopment() bool {
	return c.Env == "development"
}

// IsProd returns whether the env is prod or not
func (c *Config) IsProd() bool {
	return c.Env == "production"
}

type myenv map[string]string

func newEnv(envPath string) (myenv, error) {
	if envPath != "" {
		return godotenv.Read(envPath)
	}
	f, err := Assets.Open(".env")
	if err != nil {
		if err.Error() == "file does not exist" {
			log.Println(".env does not exist")
			return myenv{}, nil
		}
		return myenv{}, err
	}
	defer f.Close()
	return godotenv.Parse(f)
}

func (e myenv) get(name, dft string) string {
	if osenv := os.Getenv(name); osenv != "" {
		return osenv
	}
	if val, ok := e[name]; ok {
		return val
	}
	return dft
}
