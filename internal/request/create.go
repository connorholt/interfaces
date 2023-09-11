package request

type CreateTask struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

type CreateAudio struct {
	Path string `json:"path"`
	ID   string `json:"id"`
}
