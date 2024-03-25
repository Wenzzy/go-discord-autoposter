package converters

import (
	"github.com/wenzzy/go-discord-autoposter/internal/models"
	"github.com/wenzzy/go-discord-autoposter/internal/repositories/message/model"
)

func ToMessageCreateFromSendMessageService(d *models.SendMessage) model.MessageCreate {
	return model.MessageCreate{
		ChannelID:   d.ChannelID,
		Content:     d.Content,
		Attachments: d.Attachments,
	}
}
