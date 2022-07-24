package service

import (
	"github.com/giovannitgl/video-services/content-service/internal/config"
	"github.com/giovannitgl/video-services/content-service/internal/entities"
)

func VideoGetOne(id uint) *entities.Video {
	var video entities.Video
	err := config.DB.Db.First(&video, id).Error
	if err != nil {
		return nil
	}
	return &video
}

func VideoPaginatedFilter(paginator Paginator) []entities.Video {
	var videos []entities.Video
	config.DB.Db.Scopes(config.Paginate(paginator.Page, paginator.PageSize)).Find(&videos)
	return videos
}

func VideoCreate(video *entities.Video) error {
	return config.DB.Db.Create(video).Error
}

func VideoUpdate(video *entities.Video) error {
	return config.DB.Db.Save(video).Error
}
