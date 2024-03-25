package app

import "github.com/wenzzy/go-discord-autoposter/internal/utils/logger"

func (a *App) runAutoPoster() error {
	logger.Debug("AutoPoster is running")
	a.autoPoster.Run()
	return nil
}
