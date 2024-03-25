package di

import (
	"context"

	repositories "github.com/wenzzy/go-discord-autoposter/internal/repositories"
	"github.com/wenzzy/go-discord-autoposter/internal/repositories/message"
)

func (s *Container) MessageRepository(ctx context.Context) repositories.MessageRepository {
	if s.messageRepository == nil {
		s.messageRepository = message.NewRepository(s.DBClient(ctx))
	}

	return s.messageRepository
}
