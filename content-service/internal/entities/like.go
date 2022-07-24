package entities

import (
	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	VideoID uint `json:"video_id"`
	UserID  uint `json:"user_id"`
	Like    bool `json:"like" validate:"required"`
}
