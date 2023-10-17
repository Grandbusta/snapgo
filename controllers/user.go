package controllers

import (
	"net/http"

	"github.com/Grandbusta/snapgo/models"
	"github.com/Grandbusta/snapgo/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	DB *mongo.Client
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	var userInput models.UserInput
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		utils.ServerResponse(ctx, http.StatusUnprocessableEntity, "Invalid payload")
		return
	}
	if isValid := utils.ValidEmail(userInput.Email); !isValid {
		utils.ServerResponse(ctx, http.StatusUnprocessableEntity, "Invalid payload")
		return
	}
	password, err := utils.HashPassword(userInput.Password)
	if err != nil {
		utils.ServerResponse(ctx, http.StatusInternalServerError, "An error occured")
		return
	}
	user := models.User{}
	user.Password = password
	user.Email = userInput.Email

}
