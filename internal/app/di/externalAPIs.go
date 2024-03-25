package di

import (
	"context"

	"github.com/wenzzy/go-discord-autoposter/internal/external"
	"github.com/wenzzy/go-discord-autoposter/internal/external/discord"
)

func (s *Container) DiscordAPI(_ context.Context) external.Discord {
	if s.discordAPI == nil {
		discordConfig := s.DiscordConfig()
		s.discordAPI = discord.NewAPI(
			discordConfig.AccessToken(),
		)
	}

	return s.discordAPI
}
