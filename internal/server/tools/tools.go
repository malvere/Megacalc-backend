package tools

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type responseWriter struct {
	http.ResponseWriter
	code int
}

type Tools struct {
	logger logrus.Logger
}

// Constants for contexts
const (
	sessionName        = "studopolis"
	ctxKeyUser  ctxKey = iota
	ctxKeyRequesID
)

type ctxKey int8

func NewTools() *Tools {
	t := &Tools{
		logger: *logrus.New(),
	}
	return t
}

func (s *Tools) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_adr": r.RemoteAddr,
		})
		logger.Infof("Started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		logger.Infof(
			"Completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Since(start),
		)
	})
}

func (s *Tools) SetRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequesID, id)))
	})
}

func (s *Tools) Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		bytes, ok := data.([]byte)
		if !ok {
			json.NewEncoder(w).Encode(data)
		} else {
			w.Write(bytes)
		}
	}
}

func (s *Tools) Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.Respond(w, r, code, err)
}
