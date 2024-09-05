package animal

import (
	"testing"

	"github.com/sarff/prjctr-golang_beginning/gocourse05/zooobservation/camera"
	"github.com/sarff/prjctr-golang_beginning/gocourse05/zooobservation/common"
)

func TestAnimal_Move(t *testing.T) {
	direct := common.Left
	var historyItems []common.HistoryItem

	tiger := Animal{
		Id: 1,
		Camera: &camera.DayCamera{
			Screenshot: "day_screenshot.png",
		},
		Species: "tiger",
	}

	history, err := tiger.Move(direct, historyItems)
	if err != nil {
		t.Fatalf("Error saving history to server with Move() Tiger: %v", err)
	} else if history[0].Direction != direct || history[0].AnimalID != tiger.Id {
		t.Errorf("We got direction=%s and animalID=%d instead of: %s, %d", history[0].Direction,
			history[0].AnimalID, direct, tiger.Id)
	}
}

func TestDayCamera_DetectMovement(t *testing.T) {
	direct := common.Left
	var historyItems []common.HistoryItem

	tiger := Animal{
		Id: 1,
		Camera: &camera.DayCamera{
			Screenshot: "day_screenshot.png",
		},
		Species: "tiger",
	}

	history, err := tiger.Camera.DetectMovement(direct, historyItems, tiger.Id)
	if err != nil {
		t.Fatalf("Error saving history to server with DetectMovement() Tiger: %v", err)
	} else if history[0].Direction != direct || history[0].AnimalID != tiger.Id {
		t.Errorf("We got direction=%s and animalID=%d instead of: %s, %d", history[0].Direction,
			history[0].AnimalID, direct, tiger.Id)
	}
}

func TestNightCamera_DetectMovement(t *testing.T) {
	direct := common.Bottom
	var historyItems []common.HistoryItem
	bear := Animal{
		Id: 2,
		Camera: &camera.NightCamera{
			Screenshot: "night_screenshot.png",
		},
		Species: "bear",
	}

	history, err := bear.Camera.DetectMovement(direct, historyItems, bear.Id)
	if err != nil {
		t.Fatalf("Error saving history to server with DetectMovement() Bear: %v", err)
	} else if history[0].Direction != direct || history[0].AnimalID != bear.Id {
		t.Errorf("We got direction=%s and animalID=%d instead of: %s, %d", history[0].Direction,
			history[0].AnimalID, direct, bear.Id)
	}
}
