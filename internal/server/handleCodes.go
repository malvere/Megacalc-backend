package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"tg-backend/internal/db"
	"tg-backend/internal/db/sqlc"

	"github.com/google/uuid"
)

func (s *Server) handleCodes() http.HandlerFunc {
	type request struct {
		Code  string `json:"code,omitempty"`
		State bool   `json:"state,omitempty"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			log.Printf("Error during decode: %s", err)
		}

		switch r.Method {

		case http.MethodGet:
			// GET
			code := r.URL.Query().Get("code")

			c, err := createCode(code, s.store)
			if err != nil {
				s.tools.Error(w, r, http.StatusInternalServerError, err)
			}
			s.tools.Respond(w, r, http.StatusOK, &c)

		case http.MethodPost:
			// POST
			c, err := findCode(req.Code, s.store)
			log.Print(c.Active)
			if err != nil {
				s.tools.Error(w, r, http.StatusInternalServerError, err)
			}
			s.tools.Respond(w, r, http.StatusOK, &c)

		case http.MethodPatch:
			// PATCH
			err := alterCode(req.Code, s.store, req.State)
			if err != nil {
				s.tools.Error(w, r, http.StatusInternalServerError, err)
			}
			s.tools.Respond(w, r, http.StatusOK, "Code "+req.Code+" is set to "+strconv.FormatBool(req.State))

		case http.MethodDelete:
			// DELETE
			err := deleteCode(req.Code, s.store)
			if err != nil {
				s.tools.Error(w, r, http.StatusInternalServerError, err)
			}
			s.tools.Respond(w, r, http.StatusOK, "Code "+req.Code+" deleted")
		}

	}
}
func createCode(code string, store *db.Store) (sqlc.InviteCode, error) {
	arg := sqlc.CreateCodeParams{
		CodeID: uuid.New(),
		Code:   code,
		Active: true,
	}

	c, err := store.CreateCode(context.Background(), arg)
	if err != nil {
		return c, err
	}
	return c, nil
}

func deleteCode(code string, store *db.Store) error {
	err := store.DeleteCode(context.Background(), code)
	if err != nil {
		return err
	}
	return nil
}

func findCode(code string, store *db.Store) (sqlc.InviteCode, error) {
	c, err := store.FindCode(context.Background(), code)
	if err != nil {
		return c, err
	}
	return c, nil
}

func alterCode(code string, store *db.Store, newState bool) error {
	arg := sqlc.UpdateCodeParams{
		Code:   code,
		Active: newState,
	}

	err := store.UpdateCode(context.Background(), arg)
	if err != nil {
		return err
	}
	return nil
}
