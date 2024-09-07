package sensor

import (
	"math/rand/v2"
	"sync"

	"github.com/sarff/prjctr-golang_beginning/gocourse07/birdaviary/internal/logger"

	"github.com/sarff/prjctr-golang_beginning/gocourse07/birdaviary/internal/centralsystem"
)

type Sensor interface {
	Start()
	Stop()
}

type MainSensor struct {
	name string
	cs   *centralsystem.CentralSystem
	log  *logger.Logger
}

func (s *MainSensor) Start(wg *sync.WaitGroup) {
	defer wg.Done()
	s.log.Info("Starting sensor ", "sensor", s.name)
	data := s.GenerateSensorData()
	s.cs.SaveData(s.name, data)
}

func (s *MainSensor) Stop() {
	s.log.Info("Stopping sensor ", "sensor", s.name)
}

func (s *MainSensor) GenerateSensorData() int {
	return rand.IntN(100)
}
