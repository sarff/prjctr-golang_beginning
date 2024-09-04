package animal

import "zooobservation/camera"

type Animal struct {
	Id      int
	Camera  camera.Camera
	Species string
}

func (t *Animal) Move(direction camera.Direction, historyItems []camera.HistoryItem) ([]camera.HistoryItem, error) {
	return t.Camera.DetectMovement(direction, historyItems, t.Id)
}
