package cmd

import (
	"github.com/Grandbusta/snapgo/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	port := config.GetEnvs().PORT
	s.InitializeApp()
	err := s.router.Run(port)
	if err != nil {
		logrus.Fatal(err)
	}
}
