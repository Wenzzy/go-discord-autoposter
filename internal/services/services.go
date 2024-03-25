package services

import (
	"context"

	"github.com/wenzzy/go-discord-autoposter/internal/models"
)

type MessageService interface {
	Send(ctx context.Context, d models.SendMessage) error
}
