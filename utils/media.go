package utils

import (
	"context"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/admin"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func VideoUploadHelper(input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	CloudName := os.Getenv("CLOUD_NAME")
	CloudAPIKey := os.Getenv("CLOUD_API_KEY")
	CloudAPISecret := os.Getenv("CLOUD_API_SECRET")
	CloudUploadFolder := os.Getenv("CLOUD_UPLOAD_FOLDER")

	cld, err := cloudinary.NewFromParams(CloudName, CloudAPIKey, CloudAPISecret)
	if err != nil {
		return "", err
	}

	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: CloudUploadFolder})
	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}

func GetVideosHelper() (*admin.AssetsResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	CloudName := os.Getenv("CLOUD_NAME")
	CloudAPIKey := os.Getenv("CLOUD_API_KEY")
	CloudAPISecret := os.Getenv("CLOUD_API_SECRET")

	cld, err := cloudinary.NewFromParams(CloudName, CloudAPIKey, CloudAPISecret)
	if err != nil {
		return nil, err
	}

	resp, err := cld.Admin.Assets(ctx, admin.AssetsParams{})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
