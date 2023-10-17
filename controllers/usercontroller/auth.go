package usercontroller

import (
	"go-latihan/config"
	"go-latihan/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	hashPasword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPasword)
	config.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "register berhasil"})

}

func Index(c *gin.Context) {
	var users []models.User
	config.DB.Preload("Posts").Find(&users)

	c.JSON(http.StatusOK, gin.H{"users": users})

}
