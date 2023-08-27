package api

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/RianNegreiros/short-video-system/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	title := c.PostForm("title")
	description := c.PostForm("description")
	category := c.PostForm("category")
	tags := c.PostFormArray("tags")

	filename := uuid.New().String() + filepath.Ext(file.Filename)
	filepath := filepath.Join(os.Getenv("VIDEO_UPLOAD_PATH"), filename)

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user ID"})
		return
	}

	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert user ID"})
		return
	}

	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	if err := saveVideoInfoToDB(filename, filepath, title, description, category, tags, userIDUint); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save video info to database"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Video uploaded successfully"})
}

func GetVideos(c *gin.Context) {
	var videos []models.Video

	if err := models.GetDB().Find(&videos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve videos"})
		return
	}

	c.JSON(http.StatusOK, videos)
}

func saveVideoInfoToDB(filename, filepath, title, description, category string, tags []string, userID uint) error {
	video := models.Video{
		UserID:      userID,
		Title:       title,
		Description: description,
		Category:    category,
		Tags:        tags,
		FileName:    filename,
		FilePath:    filepath,
	}

	if err := models.GetDB().Create(&video).Error; err != nil {
		return err
	}

	return nil
}
