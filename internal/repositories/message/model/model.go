package model

type Message struct {
	ID          uint    `db:"id"`
	ChannelID   int64   `db:"channel_id"`
	Content     *string `db:"content"`
	Attachments *string `db:"attachments"`

	IsError bool `db:"is_error"`

	CreatedAt string `db:"created_at"`
}

type MessageCreate struct {
	ChannelID   int64    `db:"channel_id"`
	Content     *string  `db:"content"`
	Attachments []string `db:"attachments"`
}

type MessageUpdate struct {
	Content     *string  `db:"content"`
	Attachments []string `db:"attachments"`
	IsError     *bool    `db:"is_error"`
}
