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

	promoCodes := s.router.PathPrefix("/promo").Subrouter()
	promoCodes.Use(s.handlePromoToken)
	promoCodes.Handle("/code", s.handlePromoCodes())

	secure := s.router.PathPrefix("/secure").Subrouter()
	secure.Use(s.handleToken)
	secure.Handle("/code", s.handleCodes())
	secure.Handle("/list-all-codes", s.handleListAllCodes())
	secure.Handle("/list-all-users", s.handleListAllUsers())
	secure.Handle("/user", s.handleUsers())
}
