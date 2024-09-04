package camera

import (
	"fmt"
	"testing"
	"time"

	"zooobservation/common"
)

func TestDayCameraSaveToServer(t *testing.T) {
	historyItems := []common.HistoryItem{{
		Time:      time.Now(),
		Direction: common.Right,
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
	historyItems := []common.HistoryItem{{
		Time:      time.Now(),
		Direction: common.Right,
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

func TestMoveToFront(t *testing.T) {
	historyItems := []common.HistoryItem{{
		Time:      time.Now(),
		Direction: common.Right,
		AnimalID:  1,
	}}
	history := moveToFront(common.Bottom, historyItems, -1)
	if history[1].AnimalID != -1 {
		t.Errorf("history[1].AnimalID = %v, want -1", history[1].AnimalID)
	}
	if history[0].Direction != common.Bottom {
		t.Errorf("history[0].Direction = %v, want Bottom", history[0].Direction)
	}
	fmt.Println(history)
}
