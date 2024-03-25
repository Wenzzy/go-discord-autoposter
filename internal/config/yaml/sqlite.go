package yaml

import (
	"os"

	yml "gopkg.in/yaml.v3"

	"github.com/wenzzy/go-discord-autoposter/internal/config"
)

type sqliteConfig struct {
	EFile string `yaml:"db_file,required"`
}

func NewSQLiteConfig() (config.SQLiteConfig, error) {
	config := sqliteConfig{}
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

func (cfg *sqliteConfig) File() string {
	if cfg.EFile == "" {
		return "autoposter.db"
	}
	return cfg.EFile
}
