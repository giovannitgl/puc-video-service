package contract

type Event interface {
	EventName() string
}

type VideoUploadEvent struct {
	ID  uint   `json:"id"`
	Url string `json:"url"`
}

func (VideoUploadEvent) EventName() string {
	return "video.upload"
}
