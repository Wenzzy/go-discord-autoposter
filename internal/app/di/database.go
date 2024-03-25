package di

import (
	"context"
	"database/sql"
	"os"

	"github.com/pressly/goose/v3"
	"go.uber.org/zap"

	"github.com/wenzzy/go-discord-autoposter/internal/utils/closer"
	"github.com/wenzzy/go-discord-autoposter/internal/utils/logger"
	"github.com/wenzzy/go-discord-autoposter/migrations"
)

func (s *Container) DBClient(_ context.Context) *sql.DB {

	s.MessagesConfig()

	if _, err := os.Stat("autoposter.db"); err != nil {
		logger.Debug("Creating new DB")

		file, err := os.Create("autoposter.db")
		if err != nil {
			logger.Fatal("Failed to create DB", zap.Error(err))
		}

		_ = file.Close()
		logger.Debug("DB created")
	}

	db, err := sql.Open("sqlite3", "autoposter.db")
	if err != nil {
		logger.Fatal("Failed to open DB", zap.Error(err))
	}

	closer.Add(db.Close)
	// goose.SetLogger(goose.NopLogger())

	goose.SetBaseFS(migrations.EmbedMigrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "."); err != nil {
		panic(err)
	}

	return db
}
