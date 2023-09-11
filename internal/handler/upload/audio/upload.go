package audio

import (
	"encoding/json"
	"net/http"

	"github.com/connorholt/interfaces/internal/request"
	"github.com/connorholt/interfaces/internal/storage"
)

type Upload struct {
	storage *storage.InMemory
}

func NewUpload(storage *storage.InMemory) *Upload {
	return &Upload{
		storage: storage,
	}
}

func (h *Upload) Create(w http.ResponseWriter, r *http.Request) {
	var req request.CreateAudio

	data := json.NewDecoder(r.Body)
	_ = data.Decode(&req)

	//m := model.Audio{
	//	Path: req.Path,
	//}
	//
	//h.storage.Add(req.ID, m)

	w.Write([]byte("ok"))
}
