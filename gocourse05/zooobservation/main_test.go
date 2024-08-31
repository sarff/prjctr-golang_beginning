package main

import (
	"testing"
	"time"
)

func TestSaveToServer(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code have panic")
		}
	}()

	err := error(nil)
	history := &History{{
		time:      time.Now(),
		direction: right,
		animalID:  1,
	}}

	dayCamera := DayCamera{
		screenshot: "day_screenshot.png",
	}
	nightCamera := NightCamera{
		screenshot: "night_screenshot.png",
	}

	err = dayCamera.SaveToServer(history)
	if err != nil {
		t.Errorf("Error saving history to server with DayCam")
	}

	err = nightCamera.SaveToServer(history)
	if err != nil {
		t.Errorf("Error saving history to server with NightCam")
	}
}

func TestDetectMovement(t *testing.T) {
	direct := Direction(left)
	history := &History{}

	tiger := Animal{
		id: 1,
		cam: &DayCamera{
			screenshot: "day_screenshot.png",
		},
		species: "tiger",
	}
	bear := Animal{
		id: 2,
		cam: &NightCamera{
			screenshot: "night_screenshot.png",
		},
		species: "bear",
	}
	err := tiger.cam.DetectMovement(direct, history, tiger.id)
	if err != nil {
		t.Errorf("Error saving history to server with Tiger")
	}
	err = bear.cam.DetectMovement(direct, history, bear.id)
	if err != nil {
		t.Errorf("Error saving history to server with Bear")
	}
}
