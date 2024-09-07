package sensor

import (
	"github.com/sarff/prjctr-golang_beginning/gocourse07/birdaviary/internal/centralsystem"
	"github.com/sarff/prjctr-golang_beginning/gocourse07/birdaviary/internal/logger"
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
