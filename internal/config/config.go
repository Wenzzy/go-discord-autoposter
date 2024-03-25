package config

import (
	"github.com/joho/godotenv"

	"github.com/wenzzy/go-discord-autoposter/internal/utils/str2dur"
)

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}

type SQLiteConfig interface {
	File() string
}

type Message struct {
	Topic       *string                  `yaml:"topic"`
	ChannelName *string                  `yaml:"channel_name"`
	ChannelID   int64                    `yaml:"channel_id"`
	Interval    str2dur.ExtendedDuration `yaml:"interval"`
	Text        string                   `yaml:"text"`
	Files       []string                 `yaml:"files"`
}

type MessagesConfig interface {
	Messages() []Message
}

type DiscordConfig interface {
	AccessToken() string
}

type AppConfig interface {
	Env() string
}

type LoggerConfig interface {
	Level() string
	Path() string
}
