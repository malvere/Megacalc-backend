package server

import "tg-backend/internal/server/tools"

func (s *Server) configRouter() {
	s.router.Use(s.tools.SetRequestID)
	s.router.Use(s.tools.LogRequest)
	s.router.Use(tools.SetCors(s.logger).Handler)

	s.router.Handle("/ping", s.handlePing())

	// s.router.Handle("/tg", s.handleGetChatMember2())
	s.router.Handle("/user", s.handleUsers()).Methods("GET", "OPTIONS")

	users := s.router.PathPrefix("/tg").Subrouter()
	users.Use(s.handleAuth)
	users.Handle("", s.handleGetChatMember2())

	codes := s.router.PathPrefix("/code").Subrouter()
	codes.Use(s.handleToken)
	codes.Handle("/", s.handleCodes())
}
