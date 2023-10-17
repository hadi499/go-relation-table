package postcontroller

import (
	"go-latihan/config"
	"go-latihan/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var posts []models.Post
	config.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{"posts": posts})

}

func Create(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	config.DB.Create(&post)
	c.JSON(http.StatusOK, gin.H{"message": "create post berhasil"})

}
