package config

import (
	"github.com/giovannitgl/video-services/user-service/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	db.AutoMigrate(&entities.User{})
	DB = DBInstance{Db: db}
}
