package cmd

import (
	"fmt"

	"github.com/Grandbusta/snapgo/internal/routes"
)

func (s *Server) registerRoutes() {
	routes.UserRoutes(s.router)
}

func (s *Server) InitializeApp() {
	s.registerRoutes()
	fmt.Println("App initialized")
}
