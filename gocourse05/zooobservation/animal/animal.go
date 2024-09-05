package animal

import (
	"github.com/sarff/prjctr-golang_beginning/gocourse05/zooobservation/camera"
	"github.com/sarff/prjctr-golang_beginning/gocourse05/zooobservation/common"
)

type Animal struct {
	Id      int
	Camera  camera.Camera
	Species string
}

func (t *Animal) Move(direction common.Direction, historyItems []common.HistoryItem) ([]common.HistoryItem, error) {
	return t.Camera.DetectMovement(direction, historyItems, t.Id)
}
