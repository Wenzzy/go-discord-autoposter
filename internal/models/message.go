package models

import (
	"time"

	"github.com/wenzzy/go-discord-autoposter/internal/utils/str2dur"
)

type Message struct {
	ID          uint
	ChannelID   int64
	Content     *string
	Attachments []string

	IsError bool

	CreatedAt time.Time
}

type SendMessage struct {
	ChannelName *string
	Topic       *string
	ChannelID   int64
	Content     *string
	Attachments []string
	Interval    str2dur.ExtendedDuration
}
