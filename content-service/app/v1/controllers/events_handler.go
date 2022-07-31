package controllers

import (
	"fmt"
	"github.com/giovannitgl/video-services/content-service/internal/contract"
	"github.com/giovannitgl/video-services/content-service/internal/service"
)

func UpdateVideoUrl(event *contract.VideoUploadEvent) error {
	video := service.VideoGetOne(event.ID)
	if video != nil {
		return fmt.Errorf("video id %d does not exist", event.ID)
	}

	video.VideoUrl = event.Url
	return service.VideoUpdate(video)
}
