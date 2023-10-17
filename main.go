package main

import (
	"go-latihan/config"
	"go-latihan/controllers/postcontroller"
	"go-latihan/controllers/usercontroller"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()
	r.GET("/api/posts", postcontroller.Index)
	r.GET("/api/users", usercontroller.Index)
	r.POST("/api/register", usercontroller.Register)
	r.POST("/api/posts", postcontroller.Create)

	r.Run() // listen and serve on 0.0.0.0:8080
}
