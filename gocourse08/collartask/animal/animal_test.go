package animal

import (
	"math/rand/v2"
	"reflect"
	"testing"
)

func TestNewAnimal(t *testing.T) {
	validAnimals := map[string]bool{
		"bear":  true,
		"tiger": true,
		"lion":  true,
	}

	animal := New[float64](rand.IntN(100), rand.Float64(), Tiger)

	if !validAnimals[string(animal.AnimalType)] {
		t.Errorf("Expected AnimalType to be one of bear, tiger, or lion, but got %v", animal.AnimalType)
	}
}

func TestDataAnimal_RecordBreathing(t *testing.T) {
	animalData := DataAnimal[float64]{
		AnimalType:  "lion",
		Pulse:       100,
		Temperature: 39,
		Breaths:     []float64{},
	}

	breathData := 73.3

	animalData.RecordBreathing(breathData)

	wantData := []float64{breathData}
	if !reflect.DeepEqual(animalData.Breaths, wantData) {
		t.Errorf("DataAnimal.RecordBreathing() = %v, want %v", animalData.Breaths, wantData)
	}
}

func TestDataAnimal_RecordAnimal(t *testing.T) {
	animalData := DataAnimal[float64]{
		AnimalType:  "lion",
		Pulse:       100,
		Temperature: 39,
		Sounds:      []float64{},
	}

	soundData := 53.3

	animalData.RecordSound(soundData)

	wantData := []float64{soundData}
	if !reflect.DeepEqual(animalData.Sounds, wantData) {
		t.Errorf("DataAnimal.RecordSound() = %v, want %v", animalData.Sounds, wantData)
	}
}

func TestDataAnimal_Typify(t *testing.T) {
	validAnimals := map[string]bool{
		"bear":  true,
		"tiger": true,
		"lion":  true,
	}

	animal := New[float64](rand.IntN(100), rand.Float64(), Lion)
	tapify := animal.Typify()
	if !validAnimals[string(tapify)] {
		t.Errorf("Expected AnimalType to be one of bear, tiger, or lion, but got %v", tapify)
	}
}
