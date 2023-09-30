package actions

import (
	"github.com/Grandbusta/snapgo/internal/domain/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserAction struct {
	UserRepository repository.UserRepository
	db             *mongo.Client
}

func NewUserAction(
	UserRepository *repository.UserRepository,
	db *mongo.Client,
) *UserAction {
	return &UserAction{
		UserRepository: *UserRepository,
		db:             db,
	}
}

func (u *UserAction) CreateUser(c *gin.Context) {

	// data, err := u.UserRepository.CreateUser()
}
