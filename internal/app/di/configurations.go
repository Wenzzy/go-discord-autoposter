package di

import (
	"log"

	"github.com/wenzzy/go-discord-autoposter/internal/config"
	"github.com/wenzzy/go-discord-autoposter/internal/config/yaml"
)

func (s *Container) MessagesConfig() config.MessagesConfig {
	if s.messagesConfig == nil {
		cfg, err := yaml.NewMessagesConfig()
		if err != nil {
			log.Fatalf("failed to get sqlite config: %s", err.Error())
		}

		s.messagesConfig = cfg
	}

	return s.messagesConfig
}

func (s *Container) SQLiteConfig() config.SQLiteConfig {
	if s.sqliteConfig == nil {
		cfg, err := yaml.NewSQLiteConfig()
		if err != nil {
			log.Fatalf("failed to get sqlite config: %s", err.Error())
		}

		s.sqliteConfig = cfg
	}

	return s.sqliteConfig
}

func (s *Container) AppConfig() config.AppConfig {
	if s.appConfig == nil {
		cfg, err := yaml.NewAppConfig()
		if err != nil {
			log.Fatalf("failed to get app config: %s", err.Error())
		}

		s.appConfig = cfg
	}

	return s.appConfig
}

func (s *Container) LoggerConfig() config.LoggerConfig {
	if s.loggerConfig == nil {
		cfg, err := yaml.NewLoggerConfig()
		if err != nil {
			log.Fatalf("failed to get logger config: %s", err.Error())
		}

		s.loggerConfig = cfg
	}

	return s.loggerConfig
}

func (s *Container) DiscordConfig() config.DiscordConfig {
	if s.discordConfig == nil {
		cfg, err := yaml.NewDiscordConfig()
		if err != nil {
			log.Fatalf("failed to get auth config: %s", err.Error())
		}

		s.discordConfig = cfg
	}

	return s.discordConfig
}
