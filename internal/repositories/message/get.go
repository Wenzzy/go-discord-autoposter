package message

import (
	"context"

	"github.com/wenzzy/go-discord-autoposter/internal/models"
	"github.com/wenzzy/go-discord-autoposter/internal/repositories/message/converter"
	"github.com/wenzzy/go-discord-autoposter/internal/repositories/message/model"
)

func (r *repo) Get(ctx context.Context, messageID uint) (*models.Message, error) {
	var message model.Message
	sql := `
		SELECT 
			id,
			channel_id,
			content,
			attachments,
			is_error,
			created_at
		FROM messages
		WHERE id = $1
		ORDER BY id DESC
		LIMIT 1
	`

	err := r.db.QueryRowContext(ctx, sql, messageID).Scan(&message.ID, &message.ChannelID, &message.Content, &message.Attachments, &message.IsError, &message.CreatedAt)

	if err != nil {
		return nil, err
	}

	return converter.ToMessageFromMessageRepo(message)
}
