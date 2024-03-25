package env

import (
	cenv "github.com/caarlos0/env/v10"

	"github.com/wenzzy/go-discord-autoposter/internal/config"
)

type sqliteConfig struct {
	EFile string `env:"SQLITE_FILE,required"`
}

func NewSQLiteConfig() (config.SQLiteConfig, error) {
	config := sqliteConfig{}
	if err := cenv.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (cfg *sqliteConfig) File() string {
	return cfg.EFile
}
