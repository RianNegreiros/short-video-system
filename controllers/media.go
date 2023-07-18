package controllers

import (
	"net/http"

	"github.com/RianNegreiros/short-video-system/models"
	"github.com/RianNegreiros/short-video-system/service"
	"github.com/gin-gonic/gin"
)

func GetMedia() gin.HandlerFunc {
	return func(c *gin.Context) {
		media, err := service.NewMediaUpload().GetFiles()
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				models.MediaDto{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Error getting media"},
				})
			return
		}

		c.JSON(
			http.StatusOK,
			models.MediaDto{
				StatusCode: http.StatusOK,
				Message:    "success",
				Data:       map[string]interface{}{"data": media},
			})
	}
}

func FileUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		formfile, _, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				models.MediaDto{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Select a file to upload"},
				})
			return
		}

		uploadUrl, err := service.NewMediaUpload().FileUpload(models.File{File: formfile})
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				models.MediaDto{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Error uploading file"},
				})
			return
		}

		c.JSON(
			http.StatusOK,
			models.MediaDto{
				StatusCode: http.StatusOK,
				Message:    "success",
				Data:       map[string]interface{}{"data": uploadUrl},
			})
	}
}

func RemoteUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		var url models.Url

		if err := c.BindJSON(&url); err != nil {
			c.JSON(
				http.StatusBadRequest,
				models.MediaDto{
					StatusCode: http.StatusBadRequest,
					Message:    "error",
					Data:       map[string]interface{}{"data": err.Error()},
				})
			return
		}

		uploadUrl, err := service.NewMediaUpload().RemoteUpload(url)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				models.MediaDto{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Error uploading file"},
				})
			return
		}

		c.JSON(
			http.StatusOK,
			models.MediaDto{
				StatusCode: http.StatusOK,
				Message:    "success",
				Data:       map[string]interface{}{"data": uploadUrl},
			})
	}
}
