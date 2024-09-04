package animal

import (
	"testing"
	"zooobservation/camera"
)

func TestAnimal_Move(t *testing.T) {
	direct := camera.Left
	var historyItem []camera.HistoryItem

	tiger := Animal{
		Id: 1,
		Camera: &camera.DayCamera{
			Screenshot: "day_screenshot.png",
		},
		Species: "tiger",
	}

	history, err := tiger.Move(direct, historyItem)
	if err != nil {
		t.Fatalf("Error saving history to server with Move() Tiger: %v", err)
	} else if history[0].Direction != direct || history[0].AnimalID != tiger.Id {
		t.Errorf("We got direction=%s and animalID=%d instead of: %s, %d", history[0].Direction,
			history[0].AnimalID, direct, tiger.Id)
	}
}
