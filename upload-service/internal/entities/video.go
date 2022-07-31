package entities

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Path string `json:"path" validate:"required"`
}
