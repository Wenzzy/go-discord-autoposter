package yaml

import (
	"os"

	yml "gopkg.in/yaml.v3"

	"github.com/wenzzy/go-discord-autoposter/internal/config"
)

type discordConfig struct {
	EAccessToken string `yaml:"discord_token"`
}

func NewDiscordConfig() (config.DiscordConfig, error) {
	config := discordConfig{}
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

func (cfg *discordConfig) AccessToken() string {
	return cfg.EAccessToken
}
