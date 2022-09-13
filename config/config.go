package config

import (
	"sync"
	"time"

	"github.com/kelseyhightower/envconfig"
)

var (
	cfg  *Config
	once sync.Once
)

type Config struct {
	ApplicationName string `envconfig:"APP_NAME" default:"dealljobs assessment"`

	JWTExpiryDuration time.Duration `envconfig:"JWT_EXPIRY" default:"1h"`
	JWTSignatureKey   string        `envconfig:"JWT_SIGN_KEY" default:"this is a secret"`

	DBHost     string `envconfig:"DB_HOST" default:"localhost"`
	DBDriver   string `envconfig:"DB_DRIVER" default:"mysql"`
	DBUser     string `envconfig:"DB_USER" default:"root"`
	DBPassword string `envconfig:"DB_PASSWORD" default:"123456"`
	DBPort     string `envconfig:"DB_PORT" default:"3306"`
	DBName     string `envconfig:"DB_NAME" default:"dealljobs-test"`
}

func Get() *Config {
	c := Config{}
	once.Do(func() {
		envconfig.MustProcess("", &c)
		cfg = &c
	})

	return cfg
}
