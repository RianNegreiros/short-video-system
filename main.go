package main

import (
	"log"

	"github.com/RianNegreiros/short-video-system/controllers"
	"github.com/RianNegreiros/short-video-system/middlewares"
	"github.com/RianNegreiros/short-video-system/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	upload := r.Group("/api/upload")
	upload.Use(middlewares.JwtAuthMiddleware())
	upload.POST("/file", controllers.FileUpload())
	upload.POST("/url", controllers.RemoteUpload())

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")
}
