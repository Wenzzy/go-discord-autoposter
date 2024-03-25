package repositories

import (
	"context"

	serviceModels "github.com/wenzzy/go-discord-autoposter/internal/models"
	messageModel "github.com/wenzzy/go-discord-autoposter/internal/repositories/message/model"
)

type MessageRepository interface {
	Create(ctx context.Context, d messageModel.MessageCreate) (*uint, error)
	Get(ctx context.Context, id uint) (*serviceModels.Message, error)
	GetByChannelID(ctx context.Context, channelID int64) (*serviceModels.Message, error)
}
