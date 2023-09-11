package storage

import (
	"sync"

	"github.com/connorholt/interfaces/internal/model"
)

type InMemory struct {
	data map[string]model.Task
	mu   sync.RWMutex
}

func NewInMemory() *InMemory {
	return &InMemory{
		data: make(map[string]model.Task),
	}
}

func (s *InMemory) Add(key string, value model.Task) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
}

func (s *InMemory) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, key)
}

func (s *InMemory) Get(key string) model.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.data[key]
}

func (s *InMemory) GetAll() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	res := make([]string, 0, len(s.data))
	for _, row := range s.data {
		res = append(res, row.Text)
	}

	return res
}
