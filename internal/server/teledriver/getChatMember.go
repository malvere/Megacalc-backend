package teledriver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (t *Teledriver) GetChatMember(chatId string, userId string) (*TGResponse, error) {
	tglink := fmt.Sprintf("https://api.telegram.org/bot%s/getChatMember?chat_id=%s&user_id=%s", t.BotToken, chatId, userId)
	req, err := http.NewRequest(http.MethodGet, tglink, nil)
	if err != nil {
		log.Printf("Error while wrapping request to telegram API: %s", err)
		return nil, err
	}
	resp, err := t.Client.Do(req)
	if err != nil {
		log.Printf("Error while doing request to TG API: %s", err)
		return nil, err
	}
	tgResponse := &TGResponse{}
	if err = json.NewDecoder(resp.Body).Decode(tgResponse); err != nil {
		log.Printf("Error decoding json: %s", err)
		return nil, err
	}
	return tgResponse, nil
}
