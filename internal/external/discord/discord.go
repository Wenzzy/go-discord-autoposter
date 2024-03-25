package discord

type API struct {
	accessToken string
	baseURL     string
}

func NewAPI(accessToken string) *API {
	return &API{
		accessToken: accessToken,
		baseURL:     "https://discord.com/api/v9",
	}
}
