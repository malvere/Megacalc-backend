package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
)

func (s *Server) handleToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")
		log.Print(token)
		if token == os.Getenv("SECRET_TOKEN") {
			next.ServeHTTP(w, r)
		} else {
			s.tools.Error(w, r, http.StatusUnauthorized, errors.New("invalid token"))
			return
		}
	})
}

func (s *Server) handleAuth(next http.Handler) http.Handler {
	type request struct {
		UserId string `json:"user_id,omitempty"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		log.Print("Running TG check: ")
		log.Print(tg.Result)
		if tg.Result.Status == "creator" || tg.Result.Status == "administrator" || tg.Result.Status == "member" {
			log.Print("TG fired")
			s.tools.Respond(w, r, 200, tg)
			next.ServeHTTP(w, r)
		} else {
			log.Print("DB fired")
			log.Print("Running DB check: ")

			u, err := findUser(req.UserId, s.store)
			if err != nil {
				log.Print("User not found: ", err)
				s.tools.Respond(w, r, http.StatusUnauthorized, "Вы не карлик.")
				return
			}
			// Send signature if user signed via valid Code
			if err := uuid.Validate(u.InviteCodeID.String()); err != nil {
				log.Print("Error StatusBadRequest: ", tg)
				s.tools.Respond(w, r, http.StatusUnauthorized, "Вы не карлик.")
				return
			}
			next.ServeHTTP(w, r)
		}
	})
}
