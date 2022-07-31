package service

import (
	"context"
	"github.com/giovannitgl/video-services/content-service/internal/config"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
)

func UploadFile(file *multipart.FileHeader, buffer multipart.File, err error) (interface{}, error) {
	objectName := file.Filename
	fileBuffer := buffer
	contentType := file.Header["Content-Type"][0]
	fileSize := file.Size

	ctx := context.Background()
	// Upload the zip file with PutObject
	info, err := config.Minio.Client.PutObject(
		ctx,
		config.Minio.Bucket,
		objectName,
		fileBuffer,
		fileSize,
		minio.PutObjectOptions{ContentType: contentType},
	)
	return info, err
}
