package animal

import (
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
		t.Errorf("Expected AnimalType to be one of Bear, Tiger, or Lion, but got %v", animal.AnimalType)
	}
}

func TestDataAnimal_RecordBreathing(t *testing.T) {
}

func TestDataAnimal_RecordAnimal(t *testing.T) {
}

func TestDataAnimal_Typify(t *testing.T) {
}
