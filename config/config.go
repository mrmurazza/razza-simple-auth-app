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
}

func Get() *Config {
	c := Config{}
	once.Do(func() {
		envconfig.MustProcess("", &c)
		cfg = &c
	})

	return cfg
}
