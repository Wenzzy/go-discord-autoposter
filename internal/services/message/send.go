package message

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/wenzzy/go-discord-autoposter/internal/converters"
	"github.com/wenzzy/go-discord-autoposter/internal/models"
	"github.com/wenzzy/go-discord-autoposter/internal/utils/logger"
)

func (s *service) Send(ctx context.Context, d models.SendMessage) error {

	var topic, channel string

	if d.Topic != nil {
		topic = *d.Topic
	} else {
		topic = "topic-unset"
	}

	if d.ChannelName != nil {
		channel = *d.ChannelName
	} else {
		channel = strconv.FormatInt(d.ChannelID, 10)
	}

	message, err := s.messageRepository.GetByChannelID(ctx, d.ChannelID)

	if err == nil && time.Since(message.CreatedAt) <= time.Duration(d.Interval) {
		return nil
	}

	logger.Info(fmt.Sprintf("Sending message to discord, topic: %s, channel: %s", topic, channel))

	err = s.discordAPI.PostMessage(strconv.FormatInt(d.ChannelID, 10), d.Content, d.Attachments)
	if err != nil {
		return errors.Errorf("failed to send message to channel: %s, by Discord API, err: %v", channel, err)
	}

	_, err = s.messageRepository.Create(ctx, converters.ToMessageCreateFromSendMessageService(&d))
	if err != nil {
		return errors.Errorf("failed to create message in DB, with topic: %s, channel: %s, err: %v", topic, channel, err)
	}
	return nil
}
