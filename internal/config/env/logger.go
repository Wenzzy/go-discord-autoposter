package env

import (
	cenv "github.com/caarlos0/env/v10"

	"github.com/wenzzy/go-discord-autoposter/internal/config"
)

type loggerConfig struct {
	ELevel string `env:"LOG_LEVEL" envDefault:"info"`
	EPath  string `env:"LOG_PATH" envDefault:"logs/app.log"`
}

func NewLoggerConfig() (config.LoggerConfig, error) {
	config := loggerConfig{}
	if err := cenv.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (cfg *loggerConfig) Level() string {
	return cfg.ELevel
}

func (cfg *loggerConfig) Path() string {
	return cfg.EPath
}
