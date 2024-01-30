package teledriver

import "net/http"

type TGResponse struct {
	Ok     bool `json:"ok"`
	Result struct {
		User struct {
			ID           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			Username     string `json:"username"`
			LanguageCode string `json:"language_code"`
		} `json:"user"`
		Status      string `json:"status"`
		IsAnonymous bool   `json:"is_anonymous"`
	} `json:"result"`
}
type Teledriver struct {
	BotToken string
	Client   *http.Client
}

func NewTeledriver(token string) *Teledriver {
	return &Teledriver{
		BotToken: token,
		Client:   &http.Client{},
	}
}
