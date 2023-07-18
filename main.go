package main

import (
	"github.com/RianNegreiros/short-video-system/controllers"
	"github.com/RianNegreiros/short-video-system/middlewares"
	"github.com/RianNegreiros/short-video-system/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")
}
