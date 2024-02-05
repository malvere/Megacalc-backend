package server

import (
	"log"
	"net/http"
)

func (s *Server) handleGetChatMember2() http.HandlerFunc {
	type response struct {
		Key string `json:"key,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		signature, err := s.getHmacToken()
		if err != nil {
			log.Printf("Error getting HMAC key: %s", err)
			s.tools.Error(w, r, http.StatusInternalServerError, err)
		}
		s.tools.Respond(w, r, http.StatusOK, &response{Key: signature})
	}
}
