package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (s *Server) HandleGetChatMember() http.HandlerFunc {

	type request struct {
		UserId string `json:"user_id,omitempty"`
	}
	type response struct {
		Key string `json:"key,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			log.Printf("Error during decode: %s", err)
			s.tools.Error(w, r, http.StatusBadRequest, err)
		}

		tg, err := s.teledriver.GetChatMember(s.config.Telegram.ChatId, req.UserId)
		if err != nil {
			log.Printf("Error during GetChatMember: %s", err)
			s.tools.Error(w, r, http.StatusBadRequest, err)
		}
		if tg.Result.Status != "left" {
			signature, err := s.getHmacToken()
			if err != nil {
				log.Printf("Error getting HMAC key: %s", err)
				s.tools.Error(w, r, http.StatusInternalServerError, err)
			}
			s.tools.Respond(w, r, http.StatusOK, &response{Key: signature})
		} else {
			u, err := findUser(req.UserId, s.store)
			if err != nil {
				log.Print("User not found: ", err)
				s.tools.Respond(w, r, http.StatusUnauthorized, "Вы не карлик.")
			}
			// Send signature if user signed via valid Code
			if u.InviteCodeID != uuid.Nil {
				signature, err := s.getHmacToken()
				if err != nil {
					log.Printf("Error getting HMAC key: %s", err)
					s.tools.Error(w, r, http.StatusInternalServerError, err)
				}
				s.tools.Respond(w, r, http.StatusOK, &response{Key: signature})
			} else {
				log.Print("Error StatusBadRequest: ", tg)
				s.tools.Respond(w, r, http.StatusUnauthorized, "Вы не карлик.")
			}
		}
	}
}
