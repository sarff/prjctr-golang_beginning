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

	err = dayCamera.SaveToServer("left", history, 1)
	if err != nil {
		t.Errorf("Error saving history to server with DayCam")
	}

	err = nightCamera.SaveToServer("left", history, 1)
	if err != nil {
		t.Errorf("Error saving history to server with NightCam")
	}
}
