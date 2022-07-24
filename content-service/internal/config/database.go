package config

import (
	"github.com/giovannitgl/video-services/content-service/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	MaxPageSize     = 100
	DefaultPageSize = 10
)

type DBInstance struct {
	Db *gorm.DB
}

var DB DBInstance

func SetupDatabase() {
	db, err := gorm.Open(postgres.Open(PostgresDSN()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entities.Video{})
	DB = DBInstance{Db: db}
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > MaxPageSize:
			pageSize = MaxPageSize
		case pageSize <= 0:
			pageSize = DefaultPageSize
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
