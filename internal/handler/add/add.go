package add

import (
	"encoding/json"
	"net/http"

	"github.com/connorholt/interfaces/internal/model"
	"github.com/connorholt/interfaces/internal/request"
	"github.com/connorholt/interfaces/internal/storage"
)

type Add struct {
	storage *storage.InMemory
}

func NewAdd(storage *storage.InMemory) *Add {
	return &Add{
		storage: storage,
	}
}

func (h *Add) Create(w http.ResponseWriter, r *http.Request) {
	var req request.CreateTask

	data := json.NewDecoder(r.Body)
	_ = data.Decode(&req)

	m := model.Task{
		Text: req.Text,
	}

	h.storage.Add(req.Key, m)

	w.Write([]byte("ok"))
}
