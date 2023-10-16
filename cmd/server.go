package cmd

import (
	"log"

	"github.com/Grandbusta/snapgo/config"
	"github.com/Grandbusta/snapgo/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	router *gin.Engine
	db     *mongo.Client
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	config.NewLogger()
}

func NewServer() *Server {
	return &Server{
		router: gin.New(),
		db:     config.NewNoSqlDB(),
	}
}

func (s *Server) Run() {
	port := config.GetEnvs().PORT
	routes.Routes(s.router, s.db)
	err := s.router.Run(port)
	if err != nil {
		logrus.Fatal(err)
	}
}
