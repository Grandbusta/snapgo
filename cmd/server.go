package cmd

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	return &Server{
		router: gin.New(),
	}
}

func (s *Server) Run() {
	s.InitializeApp()
	s.router.Run(":8080")
}
