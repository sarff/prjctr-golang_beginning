package storage

import (
	"sync"

	"github.com/sarff/prjctr-golang_beginning/gocourse07/birdaviary/internal/logger"
)

type Storage struct {
	mu   sync.Mutex
	data map[string][]int

	log *logger.Logger
}

func NewStorage(log *logger.Logger) *Storage {
	return &Storage{
		data: make(map[string][]int),
		log:  log,
	}
}

func (m *Storage) Save(sensorName string, value int) {
	m.mu.Lock()
	m.data[sensorName] = append(m.data[sensorName], value)
	m.mu.Unlock()
	m.log.Info("Data saved:", sensorName, value)
}

func (m *Storage) Load(sensorName string) ([]int, error) {
	return m.data[sensorName], nil
}
