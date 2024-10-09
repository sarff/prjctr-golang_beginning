package main

import (
	"bytes"
	"strings"
	"sync"
	"testing"
)

func TestControlCondition(t *testing.T) {
	actualLog := &bytes.Buffer{}

	animalChan := make(chan *Animal, 1)
	wg := new(sync.WaitGroup)
	log := loggerNew(actualLog)

	animal := &Animal{
		ID:     1,
		Health: 25, // Рівень здоров'я нижче 30
		Hunger: 25, // Рівень голоду нижче 30
		Mood:   25, // Настрій нижче 30
	}
	wg.Add(1)
	go controlCondition(animalChan, wg, log)
	animalChan <- animal

	wg.Wait()
	close(animalChan)

	wantMessages := []string{
		"\"msg\":\"Animal needs help\",\"animalId\":1",
		"\"msg\":\"Animal needs to be fed\",\"animalId\":1",
		"\"msg\":\"Animal - needs to be released from the cage\",\"animalId\":1",
	}

	for _, message := range wantMessages {
		if !strings.Contains(actualLog.String(), message) {
			t.Errorf("Wanted testLog message '%s' not found", message)
		}
	}
}
