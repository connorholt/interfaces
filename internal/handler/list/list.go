package list

import (
	"encoding/json"
	"net/http"

	"github.com/connorholt/interfaces/internal/storage"
)

type List struct {
	storage *storage.InMemory
}

func NewList(storage *storage.InMemory) *List {
	return &List{
		storage: storage,
	}
}

func (h *List) GetAll(w http.ResponseWriter, r *http.Request) {
	tasks := h.storage.GetAll()

	jsonTasks, _ := json.Marshal(tasks)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonTasks)
}
