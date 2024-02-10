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

func (s *Server) handlePromoCodes() http.HandlerFunc {
	type request struct {
		PromoName   string `json:"promo_name,omitempty"`
		PromoString string `json:"promo_string,omitempty"`
		State       bool   `json:"state,omitempty"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			log.Printf("Error during decode: %s", err)
		}
		switch r.Method {
		case http.MethodGet:
			state, err := strconv.ParseBool(r.URL.Query().Get("state"))
			if err != nil {
				log.Print("Error parsing bool: ", err)
				return
			}
			pcodes, err := queryPCodes(state, s.store)
			if err != nil {
				log.Print("Error retrieving pcodes: ", err)
				return
			}
			s.tools.Respond(w, r, http.StatusOK, &pcodes)

		case http.MethodPost:
			pc, err := createPCode(req.PromoName, req.PromoString, s.store)
			if err != nil {
				s.tools.Error(w, r, http.StatusInternalServerError, err)
			}
			s.tools.Respond(w, r, http.StatusOK, &pc)

		case http.MethodDelete:
			err := deletePCode(req.PromoName, s.store)
			if err != nil {
				s.tools.Error(w, r, http.StatusInternalServerError, err)
			}
			s.tools.Respond(w, r, http.StatusOK, "Promo code deleted! Name: "+req.PromoName)

		case http.MethodPatch:
			err := alterPCode(req.PromoName, req.State, s.store)
			if err != nil {
				s.tools.Error(w, r, http.StatusInternalServerError, err)
			}
			s.tools.Respond(w, r, http.StatusOK, "Promo code altered!")
		}
	}
}

func createPCode(promoName string, promoString string, store *db.Store) (*sqlc.PromoCode, error) {
	arg := sqlc.CreatePromoCodeParams{
		PromoID:   uuid.New(),
		PromoName: promoName,
		Promo:     promoString,
		Active:    true,
	}
	u, err := store.CreatePromoCode(context.Background(), arg)
	if err != nil {
		return nil, err
	}
	return &u, nil

}

func deletePCode(promoName string, store *db.Store) error {
	err := store.DeletePromoCode(context.Background(), promoName)
	if err != nil {
		return err
	}
	return nil
}

func queryPCodes(state bool, store *db.Store) ([]sqlc.PromoCode, error) {
	u, err := store.QuerryPromoCodes(context.Background(), state)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func alterPCode(promoName string, newState bool, store *db.Store) error {
	arg := sqlc.UpdatePromoCodeParams{
		PromoName: promoName,
		Active:    newState,
	}

	err := store.UpdatePromoCode(context.Background(), arg)
	if err != nil {
		return err
	}
	return nil
}
