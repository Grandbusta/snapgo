package routes

import (
	"github.com/Grandbusta/snapgo/internal/actions"
	"github.com/Grandbusta/snapgo/internal/domain/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRoutes(g *gin.Engine, m *mongo.Client) {
	u := g.Group("/")
	user := actions.NewUserAction(
		repository.NewUserRepository(),
		m,
	)

	u.GET("/signup", user.CreateUser)
}
