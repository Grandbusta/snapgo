package routes

import (
	"github.com/Grandbusta/snapgo/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Routes(g *gin.Engine, m *mongo.Client) {
	u := g.Group("/user")
	u.GET("/signup", controllers.CreateUser)
}
