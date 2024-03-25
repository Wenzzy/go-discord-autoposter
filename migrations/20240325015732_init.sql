-- +goose Up
CREATE TABLE messages (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	channel_id NUMERIC,
	content TEXT,
	attachments TEXT,
	is_error BOOLEAN NOT NULL DEFAULT false,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS messages;