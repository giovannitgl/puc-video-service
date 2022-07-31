package main

import (
	"github.com/giovannitgl/video-services/content-service/internal/config"
	"github.com/giovannitgl/video-services/content-service/internal/upload"
	"log"
)

func main() {
	config.SetupDatabase()
	config.SetupMinio()
	config.SetupPublisher()
	app := upload.NewApplication()
	log.Fatal(app.Listen(":8000"))
}
