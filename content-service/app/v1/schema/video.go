package schema

import (
	"github.com/giovannitgl/video-services/content-service/internal/entities"
)

type VideoCreate struct {
	Title       string `json:"title" validator:"required"`
	Description string `json:"description" validator:"required"`
}

type VideoUpdate struct {
	ID          uint   `json:"id" validator:"required"`
	Title       string `json:"title" validator:"required"`
	Description string `json:"description" validator:"required"`
}

type Video struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	VideoUrl    string `json:"video_url"`
	CreatorID   uint   `json:"creator_id"`
	Published   bool   `json:"published"`
}

func VideoResponse(video entities.Video) *Video {
	return &Video{
		ID:          video.ID,
		Title:       video.Title,
		Description: video.Description,
		VideoUrl:    video.VideoUrl,
		CreatorID:   video.CreatorID,
		Published:   video.Published,
	}
}
