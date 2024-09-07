package sensor

import (
	"github.com/sarff/prjctr-golang_beginning/gocourse07/birdaviary/internal/centralsystem"
	"github.com/sarff/prjctr-golang_beginning/gocourse07/birdaviary/internal/logger"
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
