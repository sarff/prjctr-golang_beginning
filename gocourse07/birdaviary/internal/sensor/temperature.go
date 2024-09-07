package sensor

import (
	"github.com/sarff/prjctr-golang_beginning/gocourse07/birdaviary/internal/centralsystem"
	"github.com/sarff/prjctr-golang_beginning/gocourse07/birdaviary/internal/logger"
)

type TemperatureSensor struct {
	MainSensor
}

func NewTemperatureSensor(cs *centralsystem.CentralSystem, log *logger.Logger) *TemperatureSensor {
	return &TemperatureSensor{
		MainSensor{
			name: "TemperatureSensor",
			cs:   cs,
			log:  log,
		}}
}
