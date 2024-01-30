package server

import "tg-backend/internal/server/tools"

func (s *Server) configRouter() {
	s.router.Use(s.tools.SetRequestID)
	s.router.Use(s.tools.LogRequest)
	s.router.Use(tools.SetCors(s.logger).Handler)

	s.router.Handle("/hello", s.handleHW())
	s.router.Handle("/tg", s.handleGetChatMember())
}
