package animal

import (
	"testing"

	"github.com/sarff/prjctr-golang_beginning/gocourse05/zooobservation/camera"
	"github.com/sarff/prjctr-golang_beginning/gocourse05/zooobservation/common"
)

func TestAnimal_Move(t *testing.T) {
	direction := common.Left

	tiger := Animal{
		ID: 1,
		Camera: &camera.DayCamera{
			Screenshot: "day_screenshot.png",
		},
		Species: "tiger",
	}

	history, err := tiger.Move(direction, nil)
	if err != nil {
		t.Fatalf("Unexpected error: got=%v, but want=<nil>", err)
	}
	if len(history) == 0 {
		t.Fatalf("Unexpected empty history")
	}
	if history[0].ID != tiger.ID {
		t.Errorf("Wrong history animalID: got=%v, but want=%v", history[0].ID, tiger.ID)
	}
	if history[0].Direction != direction {
		t.Errorf("Wrong direction: got=%v, but want=%v", history[0].Direction, direction)
	}
}

func TestDayCamera_DetectMovement(t *testing.T) {
	direction := common.Left

	tiger := Animal{
		ID: 1,
		Camera: &camera.DayCamera{
			Screenshot: "day_screenshot.png",
		},
		Species: "tiger",
	}

	history, err := tiger.Camera.DetectMovement(direction, nil, tiger.ID)
	if err != nil {
		t.Fatalf("Unexpected error: got=%v, but want=<nil>", err)
	}
	if len(history) == 0 {
		t.Fatalf("Unexpected empty history")
	}
	if history[0].ID != tiger.ID {
		t.Errorf("Wrong history animalID: got=%v, but want=%v", history[0].ID, tiger.ID)
	}
	if history[0].Direction != direction {
		t.Errorf("Wrong direction: got=%v, but want=%v", history[0].Direction, direction)
	}
}

func TestNightCamera_DetectMovement(t *testing.T) {
	direction := common.Bottom
	var historyItems []common.HistoryItem
	bear := Animal{
		ID: 2,
		Camera: &camera.NightCamera{
			Screenshot: "night_screenshot.png",
		},
		Species: "bear",
	}

	history, err := bear.Camera.DetectMovement(direction, historyItems, bear.ID)
	if err != nil {
		t.Fatalf("Unexpected error: got=%v, but want=<nil>", err)
	}
	if len(history) == 0 {
		t.Fatalf("Unexpected empty history")
	}
	if history[0].ID != bear.ID {
		t.Errorf("Wrong history animalID: got=%v, but want=%v", history[0].ID, bear.ID)
	}
	if history[0].Direction != direction {
		t.Errorf("Wrong direction: got=%v, but want=%v", history[0].Direction, direction)
	}
}
