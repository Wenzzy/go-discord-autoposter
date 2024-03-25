package yaml

import (
	"os"

	yml "gopkg.in/yaml.v3"

	"github.com/wenzzy/go-discord-autoposter/internal/config"
)

type loggerConfig struct {
	ELevel string `yaml:"logs_level" envDefault:"info"`
	EPath  string `yaml:"logs_path" envDefault:"logs/app.log"`
}

func NewLoggerConfig() (config.LoggerConfig, error) {
	config := loggerConfig{}
	buf, err := os.ReadFile("config.yml")
	if err != nil {
		return nil, err
	}
	err = yml.Unmarshal(buf, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (cfg *loggerConfig) Level() string {
	if cfg.ELevel == "" {
		return "info"
	}
	return cfg.ELevel
}

func (cfg *loggerConfig) Path() string {
	if cfg.EPath == "" {
		return "logs/app.log"
	}
	return cfg.EPath
}
