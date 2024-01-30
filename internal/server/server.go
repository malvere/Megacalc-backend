package server

import (
	"net/http"
	"strconv"
	"tg-backend/internal/config"
	"tg-backend/internal/server/teledriver"
	"tg-backend/internal/server/tools"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	router     *mux.Router
	config     *config.Config
	tools      *tools.Tools
	logger     *logrus.Logger
	teledriver *teledriver.Teledriver
}

func newServer(cfg *config.Config) *Server {
	server := &Server{
		config:     cfg,
		router:     mux.NewRouter(),
		tools:      tools.NewTools(),
		logger:     logrus.New(),
		teledriver: teledriver.NewTeledriver(cfg.Telegram.Token),
	}
	server.configRouter()
	return server
}
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func Start(config *config.Config) error {
	s := newServer(config)
	return http.ListenAndServe(
		":"+strconv.Itoa(config.Server.Port),
		s,
	)
}
