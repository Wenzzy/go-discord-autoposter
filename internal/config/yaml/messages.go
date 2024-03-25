package yaml

import (
	"os"

	yml "gopkg.in/yaml.v3"

	"github.com/wenzzy/go-discord-autoposter/internal/config"
)

type messagesConfig struct {
	EMessages []config.Message `yaml:"messages,flow"`
}

func NewMessagesConfig() (config.MessagesConfig, error) {
	config := messagesConfig{}
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

func (cfg *messagesConfig) Messages() []config.Message {
	return cfg.EMessages
}
