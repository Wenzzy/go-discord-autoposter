package converter

import (
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/wenzzy/go-discord-autoposter/internal/models"
	"github.com/wenzzy/go-discord-autoposter/internal/repositories/message/model"
)

func ToMessageFromMessageRepo(m model.Message) (*models.Message, error) {
	var attachments []string

	if m.Attachments != nil {
		attachments = strings.Split(*m.Attachments, ",")
	}

	layout := "2006-01-02T15:04:05Z"
	createdAt, err := time.Parse(layout, m.CreatedAt)
	if err != nil {
		return nil, errors.Errorf("failed to parse created_at: %s, err: %v", m.CreatedAt, err)
	}

	return &models.Message{
		ID:          m.ID,
		ChannelID:   m.ChannelID,
		Content:     m.Content,
		Attachments: attachments,
		IsError:     m.IsError,
		CreatedAt:   createdAt,
	}, nil
}
