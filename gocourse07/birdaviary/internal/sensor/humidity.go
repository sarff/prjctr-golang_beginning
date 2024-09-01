package sensor

import (
	"birdaviary/internal/centralsystem"
	"birdaviary/internal/logger"
)

type HumiditySensor struct {
	MainSensor
}

func NewHumiditySensor(cs *centralsystem.CentralSystem, log *logger.Logger) *HumiditySensor {
	return &HumiditySensor{
		MainSensor{
			name: "HumiditySensor",
			cs:   cs,
			log:  log,
		}}
}
