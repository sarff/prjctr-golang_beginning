package centralsystem

import (
	"birdaviary/internal/logger"
	"birdaviary/internal/storage"
	"sync"
)

type CentralSystem struct {
	Memory *storage.Storage
	log    *logger.Logger
}

func NewCentralSystem(db *storage.Storage, log *logger.Logger) *CentralSystem {
	return &CentralSystem{
		Memory: db,
		log:    log,
	}
}

func (c *CentralSystem) Start(wg *sync.WaitGroup) {
	defer wg.Done()
	c.log.Info("Starting CentralSystem")
}

func (c *CentralSystem) Stop() {
	c.log.Info("Stopping CentralSystem")
}

func (c *CentralSystem) SaveData(name string, value int) {
	c.Memory.Save(name, value)
}
