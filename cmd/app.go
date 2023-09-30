package cmd

import (
	"fmt"
	"log"

	"github.com/Grandbusta/snapgo/config"
	"github.com/Grandbusta/snapgo/internal/routes"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	config.NewLogger()
}

func (s *Server) registerRoutes() {
	routes.UserRoutes(s.router, s.db)
}

func (s *Server) InitializeApp() {
	s.registerRoutes()
	fmt.Println("App initialized")
}
