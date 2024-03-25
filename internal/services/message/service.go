package message

import (
	"github.com/wenzzy/go-discord-autoposter/internal/external"
	repositories "github.com/wenzzy/go-discord-autoposter/internal/repositories"
	services "github.com/wenzzy/go-discord-autoposter/internal/services"
)

type service struct {
	discordAPI        external.Discord
	messageRepository repositories.MessageRepository
}

func NewService(
	discordAPI external.Discord,
	messageRepo repositories.MessageRepository,
) services.MessageService {
	return &service{
		discordAPI:        discordAPI,
		messageRepository: messageRepo,
	}
}

func NewMockService(deps ...interface{}) services.MessageService {
	srv := service{}

	for _, v := range deps {
		switch s := v.(type) {
		case repositories.MessageRepository:
			srv.messageRepository = s
		case external.Discord:
			srv.discordAPI = s
		}
	}

	return &srv
}
