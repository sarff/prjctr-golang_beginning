package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDayCameraSaveToServer(t *testing.T) {
	t.Run("panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Should have panic")
			}
		}()
		var s *string
		fmt.Println(*s)
	})

	var err error
	history := &[]HistoryItem{{
		time:      time.Now(),
		direction: right,
		animalID:  1,
	}}

	dayCamera := DayCamera{
		screenshot: "day_screenshot.png",
	}

	err = dayCamera.SaveToServer(*history)
	if err != nil {
		t.Errorf("Error saving history to server with DayCam: %v", err)
	}
}

func TestNightCameraSaveToServer(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code have panic")
		}
	}()

	var err error
	history := &[]HistoryItem{{
		time:      time.Now(),
		direction: right,
		animalID:  1,
	}}

	nightCamera := NightCamera{
		screenshot: "night_screenshot.png",
	}

	err = nightCamera.SaveToServer(*history)
	if err != nil {
		t.Errorf("Error saving history to server with NightCam: %v", err)
	}
}

func TestDayCamera_DetectMovement(t *testing.T) {
	direct := left
	var history []HistoryItem

	tiger := Animal{
		id: 1,
		camera: &DayCamera{
			screenshot: "day_screenshot.png",
		},
		species: "tiger",
	}

	var err error
	history, err = tiger.camera.DetectMovement(direct, history, tiger.id)
	if err != nil {
		t.Errorf("Error saving history to server with DetectMovement() Tiger: %v", err)
	} else if history[0].direction != direct || history[0].animalID != tiger.id {
		t.Errorf("We got direction=%s and animalID=%d instead of: %s, %d", history[0].direction,
			history[0].animalID, direct, tiger.id)
	}
}

func TestNightCamera_DetectMovement(t *testing.T) {
	direct := bottom
	var history []HistoryItem
	bear := Animal{
		id: 2,
		camera: &NightCamera{
			screenshot: "night_screenshot.png",
		},
		species: "bear",
	}
	var err error
	history, err = bear.camera.DetectMovement(direct, history, bear.id)
	if err != nil {
		t.Errorf("Error saving history to server with DetectMovement() Bear: %v", err)
	} else if history[0].direction != direct || history[0].animalID != bear.id {
		t.Errorf("We got direction=%s and animalID=%d instead of: %s, %d", history[0].direction,
			history[0].animalID, direct, bear.id)
	}
}

func TestAnimal_Move(t *testing.T) {
	direct := left
	var history []HistoryItem

	tiger := Animal{
		id: 1,
		camera: &DayCamera{
			screenshot: "day_screenshot.png",
		},
		species: "tiger",
	}
	var err error
	history, err = tiger.Move(direct, history)
	if err != nil {
		t.Errorf("Error saving history to server with Move() Tiger: %v", err)
	} else if history[0].direction != direct || history[0].animalID != tiger.id {
		t.Errorf("We got direction=%s and animalID=%d instead of: %s, %d", history[0].direction,
			history[0].animalID, direct, tiger.id)
	}
}
