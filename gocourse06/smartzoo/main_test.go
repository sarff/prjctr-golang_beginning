package main

import (
	"os"
	"strings"
	"sync"
	"testing"
)

func TestControlCondition(t *testing.T) {
	logFile, err := os.CreateTemp(t.TempDir(), "testLog-*.log")
	if err != nil {
		t.Fatalf("Failed to create testLog file: %v", err)
	}
	// defer os.Remove(logFile.Name())

	animalChan := make(chan *Animal, 1)
	wg := new(sync.WaitGroup)
	log := loggerNew(logFile)

	wg.Add(1)

	animal := &Animal{
		ID:     1,
		Health: 25, // Рівень здоров'я нижче 30
		Hunger: 25, // Рівень голоду нижче 30
		Mood:   25, // Настрій нижче 30
	}

	go controlCondition(animalChan, wg, log)
	animalChan <- animal

	wg.Wait()
	close(animalChan)

	err = logFile.Close()
	if err != nil {
		return
	}

	content, err := os.ReadFile(logFile.Name())
	if err != nil {
		t.Fatalf("Failed to read testLog file: %v", err)
	}

	logOutput := string(content)

	wantMessages := []string{
		"Animal with ID 1 - needs help",
		"Animal with ID 1 - needs to be fed",
		"Animal with ID 1 - needs to be released from the cage",
	}

	for _, message := range wantMessages {
		if !strings.Contains(logOutput, message) {
			t.Errorf("Wanted testLog message '%s' not found", message)
		}
	}
}
