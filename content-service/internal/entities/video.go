package entities

import (
	"gorm.io/gorm"
	"time"
)

type Video struct {
	gorm.Model
	Title        string    `json:"title" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	CreationTime time.Time `json:"birth_date" validate:"required"`
	VideoUrl     string    `json:"video_url"`
	CreatorID    uint      `json:"creator_id"`
	Published    bool      `json:"published" gorm:"default:false"`
}
