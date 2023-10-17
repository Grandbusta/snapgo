package routes

import (
	"github.com/Grandbusta/snapgo/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Routes(g *gin.Engine, db *mongo.Client) {
	u := g.Group("/user")
	userController := controllers.UserController{DB: db}

	u.GET("/signup", userController.CreateUser)
}
