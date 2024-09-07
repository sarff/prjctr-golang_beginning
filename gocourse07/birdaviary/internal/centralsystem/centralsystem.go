package centralsystem

import (
	"github.com/sarff/prjctr-golang_beginning/gocourse07/birdaviary/internal/logger"
	"github.com/sarff/prjctr-golang_beginning/gocourse07/birdaviary/internal/storage"
	"sync"
)

type CentralSystem struct {
	memory *storage.Storage
	log    *logger.Logger
}

func NewCentralSystem(db *storage.Storage, log *logger.Logger) *CentralSystem {
	return &CentralSystem{
		memory: db,
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
	c.memory.Save(name, value)
}
