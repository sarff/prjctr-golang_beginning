package sensor

import (
	"birdaviary/internal/centralsystem"
	"birdaviary/internal/logger"
)

type BrightnessSensor struct {
	MainSensor
}

func NewBrightnessSensor(cs *centralsystem.CentralSystem, log *logger.Logger) *BrightnessSensor {
	return &BrightnessSensor{
		MainSensor{
			name: "BrightnessSensor",
			cs:   cs,
			log:  log,
		}}
}
