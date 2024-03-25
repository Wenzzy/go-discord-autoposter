package yaml

import (
	"os"

	yml "gopkg.in/yaml.v3"

	"github.com/wenzzy/go-discord-autoposter/internal/config"
)

type appConfig struct {
	AppEnv string `yaml:"app_env"`
}

func NewAppConfig() (config.AppConfig, error) {
	config := appConfig{}
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

func (cfg *appConfig) Env() string {
	if cfg.AppEnv == "" {
		return "prod"
	}
	return cfg.AppEnv
}
