package server

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"tg-backend/internal/config"
	"tg-backend/internal/db"
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
	store      *db.Store
}

func newServer(cfg *config.Config, database *sql.DB) *Server {
	server := &Server{
		config:     cfg,
		router:     mux.NewRouter(),
		tools:      tools.NewTools(),
		logger:     logrus.New(),
		teledriver: teledriver.NewTeledriver(cfg.Telegram.Token),
		store:      db.New(database),
	}
	server.configRouter()
	return server
}

func newDB(cfgD *config.DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open(cfgD.Driver, cfgD.URL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func Start(config *config.Config) error {
	db, err := newDB(&config.Database)
	if err != nil {
		log.Fatal("Error connecting to db: ", err)
	}
	defer db.Close()

	s := newServer(config, db)
	return http.ListenAndServe(
		":"+strconv.Itoa(config.Server.Port),
		s,
	)
}
