package di

import (
	"context"

	services "github.com/wenzzy/go-discord-autoposter/internal/services"
	"github.com/wenzzy/go-discord-autoposter/internal/services/message"
)

func (s *Container) MessageService(ctx context.Context) services.MessageService {
	if s.messageService == nil {
		// authConfig := s.AuthConfig()
		s.messageService = message.NewService(
			s.DiscordAPI(ctx),
			s.MessageRepository(ctx),
		)
	}

	return s.messageService
}
