package message

import (
	"database/sql"

	repositories "github.com/wenzzy/go-discord-autoposter/internal/repositories"
)

type repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) repositories.MessageRepository {
	return &repo{db: db}
}
