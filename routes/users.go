package routes

import (
	"EventBooking/models"
	"EventBooking/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data"})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create the user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created!", "user": user})
}

func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data"})
		return
	}

	err = user.ValidateUser()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})

}
