package env

import (
	cenv "github.com/caarlos0/env/v10"

	"github.com/wenzzy/go-discord-autoposter/internal/config"
)

type discordConfig struct {
	EAccessToken string `env:"DISCORD_ACCESS_TOKEN,required,unset"`
	EServerID    string `env:"DISCORD_SERVER_ID,required"`
	EChannelID   string `env:"DISCORD_CHANNEL_ID,required"`
}

func NewDiscordConfig() (config.DiscordConfig, error) {
	config := discordConfig{}
	if err := cenv.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (cfg *discordConfig) AccessToken() string {
	return cfg.EAccessToken
}

func (cfg *discordConfig) ServerID() string {
	return cfg.EServerID
}

func (cfg *discordConfig) ChannelID() string {
	return cfg.EChannelID
}
