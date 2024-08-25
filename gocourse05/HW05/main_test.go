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
	testHistory := &History{{
		time:      time.Now(),
		direction: right,
		animalID:  1,
	}}

	testDayCam := DayCamera{
		screenshot: "day_screenshot.png",
	}
	testNightCam := NightCamera{
		screenshot: "night_screenshot.png",
	}

	err = testDayCam.SaveToServer("left", testHistory, 1)
	if err != nil {
		t.Errorf("Error saving history to server with DayCam")
	}

	err = testNightCam.SaveToServer("left", testHistory, 1)
	if err != nil {
		t.Errorf("Error saving history to server with NightCam")
	}
}
