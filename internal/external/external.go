package external

type Discord interface {
	PostMessage(channelID string, content *string, filesPaths []string) error
}
