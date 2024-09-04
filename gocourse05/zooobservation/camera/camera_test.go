package camera

import (
	"testing"
	"time"
	"zooobservation/animal"
)

func TestDayCameraSaveToServer(t *testing.T) {
	historyItems := []HistoryItem{{
		Time:      time.Now(),
		Direction: Right,
		AnimalID:  1,
	}}

	dayCamera := DayCamera{
		Screenshot: "day_screenshot.png",
	}

	err := dayCamera.SaveToServer(historyItems)
	if err != nil {
		t.Fatalf("Error saving history to server with DayCam: %v", err)
	}
}

func TestNightCameraSaveToServer(t *testing.T) {
	historyItems := []HistoryItem{{
		Time:      time.Now(),
		Direction: Right,
		AnimalID:  1,
	}}

	nightCamera := NightCamera{
		Screenshot: "night_screenshot.png",
	}

	err := nightCamera.SaveToServer(historyItems)
	if err != nil {
		t.Fatalf("Error saving history to server with NightCam: %v", err)
	}
}

func TestDayCamera_DetectMovement(t *testing.T) {
	direct := Left
	var historyItems []HistoryItem

	tiger := animal.Animal{
		Id: 1,
		Camera: &DayCamera{
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
	direct := Bottom
	var historyItems []HistoryItem
	bear := animal.Animal{
		Id: 2,
		Camera: &NightCamera{
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

//func TestMoveToFront(t *testing.T) {
//	direct := Left
//	var historyItems []HistoryItem
//
//	history := moveToFront(direct, historyItems, 1)
//	fmt.Println(history)
//}
