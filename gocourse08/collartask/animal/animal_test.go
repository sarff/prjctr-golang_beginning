package animal

import (
	"reflect"
	"testing"
)

func TestNewAnimal(t *testing.T) {
	validAnimals := map[string]bool{
		"bear":  true,
		"tiger": true,
		"lion":  true,
	}

	animal := NewAnimal()

	if !validAnimals[animal.AnimalType] {
		t.Errorf("Expected AnimalType to be one of bear, tiger, or lion, but got %v", animal.AnimalType)
	}
}

func TestDataAnimal_RecordBreathing(t *testing.T) {
	animalData := DataAnimal[float64]{
		AnimalType:  "lion",
		Pulse:       100,
		Temperature: 39,
		Breathing:   []float64{},
	}

	breathData := 73.3

	animalData.RecordBreathing(breathData)

	wantData := []float64{breathData}
	if !reflect.DeepEqual(animalData.Breathing, wantData) {
		t.Errorf("DataAnimal.RecordBreathing() = %v, want %v", animalData.Breathing, wantData)
	}
}

func TestDataAnimal_RecordAnimal(t *testing.T) {
	animalData := DataAnimal[float64]{
		AnimalType:  "lion",
		Pulse:       100,
		Temperature: 39,
		Sounding:    []float64{},
	}

	soundData := 53.3

	animalData.RecordSound(soundData)

	wantData := []float64{soundData}
	if !reflect.DeepEqual(animalData.Sounding, wantData) {
		t.Errorf("DataAnimal.RecordSound() = %v, want %v", animalData.Sounding, wantData)
	}
}

func TestDataAnimal_Typify(t *testing.T) {
	validAnimals := map[string]bool{
		"bear":  true,
		"tiger": true,
		"lion":  true,
	}

	animal := NewAnimal()
	tapify := animal.Typify()
	if !validAnimals[tapify] {
		t.Errorf("Expected AnimalType to be one of bear, tiger, or lion, but got %v", tapify)
	}
}
