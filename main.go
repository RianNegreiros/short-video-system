package main

import (
	"github.com/RianNegreiros/short-video-system/controllers"
	"github.com/RianNegreiros/short-video-system/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)

	r.Run(":8080")
}
