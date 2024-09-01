package sensor

import (
	"birdaviary/internal/centralsystem"
	"birdaviary/internal/logger"
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
