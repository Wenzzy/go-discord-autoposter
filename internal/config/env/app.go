package env

import (
	cenv "github.com/caarlos0/env/v10"

	"github.com/wenzzy/go-discord-autoposter/internal/config"
)

type appConfig struct {
	EAppEnv  string `env:"APP_ENV" envDefault:"prod"`
	EAppName string `env:"APP_NAME" envDefault:"unnamed-app"`
}

func NewAppConfig() (config.AppConfig, error) {

	config := appConfig{}
	if err := cenv.Parse(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (cfg *appConfig) Env() string {
	return cfg.EAppEnv
}

func (cfg *appConfig) Name() string {
	return cfg.EAppName
}
