package main

import (
	"github.com/giovannitgl/video-services/content-service/internal/config"
	"github.com/giovannitgl/video-services/content-service/internal/content"
	"log"
)

func main() {
	config.SetupDatabase()
	eventProcessor, err := content.NewEventProcessor()
	if err != nil {
		log.Panic(err)
	}
	go eventProcessor.ProcessEvent()

	app := content.NewApplication()
	log.Fatal(app.Listen(":8000"))
}
