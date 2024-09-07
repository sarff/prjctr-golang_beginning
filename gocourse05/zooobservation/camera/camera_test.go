package camera

import (
	"testing"
	"time"

	"github.com/sarff/prjctr-golang_beginning/gocourse05/zooobservation/animal"
)

func TestDayCameraSaveToServer(t *testing.T) {
	historyItems := []HistoryItem{{
		Time:      time.Now(),
		Direction: Right,
		ID:        1,
	}}

	dayCamera := DayCamera{}

	err := dayCamera.SaveToServer(historyItems)
	if err != nil {
		t.Fatalf("Error saving history to server with DayCam: %v", err)
	}
}

func TestNightCameraSaveToServer(t *testing.T) {
	historyItems := []HistoryItem{{
		Time:      time.Now(),
		Direction: Top,
		ID:        1,
	}}

	nightCamera := NightCamera{}

	err := nightCamera.SaveToServer(historyItems)
	if err != nil {
		t.Fatalf("Error saving history to server with NightCam: %v", err)
	}
}

func TestMoveToFront(t *testing.T) {
	historyItems := []HistoryItem{{
		Time:      time.Now(),
		Direction: Right,
		ID:        1,
	}}
	history := moveToFront(Bottom, historyItems, -1)
	if history[1].ID != -1 {
		t.Errorf("history[1].ID = %v, want -1", history[1].ID)
	}
	if history[0].Direction != Bottom {
		t.Errorf("history[0].Direction = %v, want Bottom", history[0].Direction)
	}
}

func Test_Move(t *testing.T) {
	direction := Left

	tiger := animal.Animal{
		ID:      1,
		Species: "tiger",
	}

	controller := Controller{
		DayCamera:   DayCamera{},
		NightCamera: NightCamera{},
	}

	history, err := controller.Move(tiger, direction, nil)
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
	direction := Left

	controller := Controller{
		DayCamera: DayCamera{},
	}

	tiger := animal.Animal{
		ID:      1,
		Species: "tiger",
	}

	history, err := controller.DayCamera.DetectMovement(direction, nil, tiger.ID)
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
	direction := Bottom

	controller := Controller{
		NightCamera: NightCamera{},
	}

	bear := animal.Animal{
		ID:      2,
		Species: "bear",
	}

	history, err := controller.NightCamera.DetectMovement(direction, nil, bear.ID)
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
