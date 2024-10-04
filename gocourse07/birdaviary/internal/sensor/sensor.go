package sensor

import (
	"math/rand/v2"
	"sync"

	"github.com/sarff/prjctr-golang_beginning/gocourse07/birdaviary/internal/logger"

	"github.com/sarff/prjctr-golang_beginning/gocourse07/birdaviary/internal/centralsystem"
)

type Sensor interface {
	Start(*sync.WaitGroup)
	Stop(*sync.WaitGroup)
}

type MainSensor struct {
	name string
	cs   *centralsystem.CentralSystem
	log  *logger.Logger
}

func (s *MainSensor) Start(wg *sync.WaitGroup) {
	defer s.Stop(wg)
	s.log.Info("Starting sensor ", "sensor", s.name)
	for range 10 {
		data := s.generateSensorData()
		s.cs.SaveData(s.name, data)
	}
}

func (s *MainSensor) Stop(wg *sync.WaitGroup) {
	defer wg.Done()
	s.log.Info("Stopping sensor", "sensor", s.name)
}

func (s *MainSensor) generateSensorData() int {
	return rand.IntN(100)
}

func NewSensor(cs *centralsystem.CentralSystem, log *logger.Logger, name string) *MainSensor {
	return &MainSensor{name, cs, log}
}
