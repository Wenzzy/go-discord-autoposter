package message

import (
	"context"
	"strings"

	"github.com/wenzzy/go-discord-autoposter/internal/repositories/message/model"
)

func (r *repo) Create(ctx context.Context, d model.MessageCreate) (*uint, error) {
	var insertedID uint
	sql := `
		INSERT INTO messages (
			channel_id, content, attachments
		) 
		VALUES (
			$1, $2, $3
		)
		RETURNING id
	`

	var attachments *string

	if d.Attachments != nil {
		joinedAttachments := strings.Join(d.Attachments, ",")
		attachments = &joinedAttachments
	}

	err := r.db.QueryRowContext(ctx, sql, d.ChannelID, d.Content, attachments).Scan(&insertedID)
	if err != nil {
		return nil, err
	}

	return &insertedID, nil
}
