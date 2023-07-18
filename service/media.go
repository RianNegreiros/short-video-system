package service

import (
	"github.com/RianNegreiros/short-video-system/models"
	"github.com/RianNegreiros/short-video-system/utils"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

type mediaUpload interface {
	FileUpload(file models.File) (string, error)
	RemoteUpload(url models.Url) (string, error)
}

type media struct{}

func NewMediaUpload() mediaUpload {
	return &media{}
}

func (*media) FileUpload(file models.File) (string, error) {
	err := validate.Struct(file)
	if err != nil {
		return "", err
	}

	uploadUrl, err := utils.VideoUploadHelper(file.File)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}

func (*media) RemoteUpload(url models.Url) (string, error) {
	err := validate.Struct(url)
	if err != nil {
		return "", err
	}

	uploadUrl, errUrl := utils.VideoUploadHelper(url.Url)
	if errUrl != nil {
		return "", err
	}
	return uploadUrl, nil
}
