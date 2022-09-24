package main

import (
	"github.com/giovannitgl/video-services/user-service/internal/user"
	"log"
)

func main() {
	app := user.NewApplication()
	log.Fatal(app.Listen(":8000"))
}
