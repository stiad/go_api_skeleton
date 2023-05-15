package app

func (s *Server) routes() {
	s.registerRoute("/", s.handleHelloWorld(), []string{"GET"})
	s.registerRoute("/protected", s.handleProtectedHelloWorld(), []string{"GET"}, s.middlewareApiKeyAuth)
}
