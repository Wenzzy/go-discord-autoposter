package di

import (
	"github.com/wenzzy/go-discord-autoposter/internal/config"
	"github.com/wenzzy/go-discord-autoposter/internal/external"
	repositories "github.com/wenzzy/go-discord-autoposter/internal/repositories"
	services "github.com/wenzzy/go-discord-autoposter/internal/services"
)

type Container struct {
	appConfig      config.AppConfig
	discordConfig  config.DiscordConfig
	sqliteConfig   config.SQLiteConfig
	loggerConfig   config.LoggerConfig
	messagesConfig config.MessagesConfig

	// redisClient redis.Client

	// Repositories =============

	messageRepository repositories.MessageRepository

	// ==========================

	// Services =================

	messageService services.MessageService

	// ==========================

	// External APIs =================

	discordAPI external.Discord

	// ==========================

}

func NewContainer() *Container {
	return &Container{}
}
