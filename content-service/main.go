package main

import (
	"github.com/giovannitgl/video-services/content-service/internal/content"
	"log"
)

func main() {
	app := content.NewApplication()
	log.Fatal(app.Listen(":8000"))
}
