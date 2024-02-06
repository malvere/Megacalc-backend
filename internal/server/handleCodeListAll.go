package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) handleListAllCodes() http.HandlerFunc {
	type request struct {
		Page int32 `json:"page,omitempty"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			log.Printf("Error during decode: %s", err)
			req.Page = 0
		}
		codes, err := s.store.ListAllCodes(context.Background(), req.Page*20)
		if err != nil {
			log.Printf("Error during decode: %s", err)
			s.tools.Error(w, r, http.StatusInternalServerError, err)
		}
		s.tools.Respond(w, r, http.StatusOK, &codes)
	}
}
