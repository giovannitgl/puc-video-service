package service

import (
	"github.com/giovannitgl/video-services/content-service/internal/config"
	"github.com/giovannitgl/video-services/content-service/internal/contract"
)

func SendVideoUploadedEvent(id uint, url string) error {
	event := contract.VideoUploadEvent{
		ID:  id,
		Url: url,
	}
	return config.Pub.Publish(event)
}
