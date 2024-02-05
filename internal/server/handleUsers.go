package server

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"tg-backend/internal/db"
	"tg-backend/internal/db/sqlc"

	"github.com/google/uuid"
)

func (s *Server) handleUsers() http.HandlerFunc {
	type request struct {
		TelegramID string `json:"telegram_id,omitempty"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			log.Printf("Error during decode: %s", err)
		}
		switch r.Method {
		case http.MethodGet:
			tgid := r.URL.Query().Get("tgid")
			code := r.URL.Query().Get("code")
			inviteCode, err := findCode(code, s.store)
			if err != nil {
				log.Print(err)
				s.tools.Error(w, r, http.StatusUnauthorized, err)
				return
			}
			u, err := createUser(tgid, s.store, &inviteCode)
			if err != nil {
				log.Print(err)
				s.tools.Error(w, r, http.StatusInternalServerError, err)
				return
			}
			s.tools.Respond(w, r, http.StatusOK, &u)
		case http.MethodPost:
			u, err := findUser(req.TelegramID, s.store)
			if err != nil {
				s.tools.Error(w, r, http.StatusInternalServerError, err)
			}
			s.tools.Respond(w, r, http.StatusOK, &u)
		case http.MethodDelete:
			err := deleteUser(req.TelegramID, s.store)
			if err != nil {
				s.tools.Error(w, r, http.StatusInternalServerError, err)
			}
			s.tools.Respond(w, r, http.StatusOK, "User deleted! Telegram ID: "+req.TelegramID)
		case http.MethodPatch:
			err := alterUser(req.TelegramID, uuid.New(), s.store)
			if err != nil {
				s.tools.Error(w, r, http.StatusInternalServerError, err)
			}
			s.tools.Respond(w, r, http.StatusOK, "User altered!")
		}
	}
}

func createUser(tgID string, store *db.Store, code *sqlc.InviteCode) (*sqlc.User, error) {
	arg := sqlc.CreateUserParams{
		UserID:       uuid.New(),
		TelegramID:   tgID,
		InviteCodeID: code.CodeID,
	}
	if code.Active {
		u, err := store.CreateUser(context.Background(), arg)
		if err != nil {
			return &u, err
		}

		err = alterCode(code.Code, store, false)
		if err != nil {
			return &u, err
		}
		return &u, nil
	} else {
		return nil, errors.New("code is activated")
	}

}

func deleteUser(tgID string, store *db.Store) error {
	err := store.DeleteByTGID(context.Background(), tgID)
	if err != nil {
		return err
	}
	return nil
}

func findUser(tgID string, store *db.Store) (sqlc.User, error) {
	u, err := store.FindByTGID(context.Background(), tgID)
	if err != nil {
		return u, err
	}
	return u, nil
}

func alterUser(tgID string, code uuid.UUID, store *db.Store) error {
	arg := sqlc.UpdateInviteByTGIDParams{
		TelegramID:   tgID,
		InviteCodeID: code,
	}

	err := store.UpdateInviteByTGID(context.Background(), arg)
	if err != nil {
		return err
	}
	return nil
}
