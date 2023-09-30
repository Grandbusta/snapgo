package actions

import (
	"github.com/Grandbusta/snapgo/internal/domain/repository"
	"github.com/gin-gonic/gin"
)

type UserAction struct {
	UserRepository repository.UserRepository
}

func NewUserAction(
	UserRepository repository.UserRepository,
) *UserAction {
	return &UserAction{
		UserRepository: UserRepository,
	}
}

func (u *UserAction) CreateUser(c *gin.Context) {

}
