package storage

import (
	"sync"

	"birdaviary/internal/logger"
)

type Storage struct {
	data map[string][]int
	mu   sync.Mutex
	log  *logger.Logger
}

func NewStorage(log *logger.Logger) *Storage {
	return &Storage{
		data: make(map[string][]int),
		log:  log,
	}
}

func (m *Storage) Save(sensorName string, value int) {
	defer m.mu.Unlock()
	m.mu.Lock()
	m.data[sensorName] = append(m.data[sensorName], value)
	m.log.Info("Data saved:", sensorName, value)
}

func (m *Storage) Load(sensorName string) ([]int, error) {
	return m.data[sensorName], nil
}
