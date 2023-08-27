package main

import (
	"github.com/RianNegreiros/short-video-system/api"
	"github.com/RianNegreiros/short-video-system/models"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := gin.Default()

	auth := router.Group("/api/auth")
	{
		auth.POST("/signup", api.SignUp)
		auth.POST("/login", api.Login)
	}

	videos := router.Group("/api/videos")
	{
		videos.GET("/", api.GetVideos)
		videos.POST("/", api.JWTMiddleware(), api.UploadFile)
	}

	router.Run(":8080")
}
