package cmd

import (
	"github.com/Grandbusta/snapgo/config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	router *gin.Engine
	db     *mongo.Client
}

func NewServer() *Server {
	return &Server{
		router: gin.New(),
		db:     config.NewNoSqlDB(),
	}
}

func (s *Server) Run() {
	s.InitializeApp()
	s.router.Run(":3000")
}
