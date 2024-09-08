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
