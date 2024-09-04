package animal

import (
	"zooobservation/camera"
	"zooobservation/common"
)

type Animal struct {
	Id      int
	Camera  camera.Camera
	Species string
}

func (t *Animal) Move(direction common.Direction, historyItems []common.HistoryItem) ([]common.HistoryItem, error) {
	return t.Camera.DetectMovement(direction, historyItems, t.Id)
}
