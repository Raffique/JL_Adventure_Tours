package controllers

import (
	"net/http"
	"server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SignUp(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    input.Password = utils.HashPassword(input.Password)
    if err := models.DB.Create(&input).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Account created"})
}

func Login(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := models.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
            return
        }
    }

    if !utils.CheckPasswordHash(input.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
        return
    }

    token, err := utils.GenerateToken(user.ID, user.Role)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"token": token})
}
