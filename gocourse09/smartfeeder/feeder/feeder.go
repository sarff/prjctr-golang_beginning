package feeder

import (
	"fmt"

	"github.com/sarff/prjctr-golang_beginning/gocourse09/smartfeeder/animal"
)

type TypeFood string

const (
	Grass       TypeFood = "Grass"
	Meat        TypeFood = "Meat"
	Berries     TypeFood = "Berries"
	GeneralFood TypeFood = "General Food"
)

type Feeder struct {
	FoodStock          map[TypeFood]int
	AnimalTypeStrategy AnimalTypeStrategy
}

type AnimalTypeStrategy interface {
	FoodType(AnimalType animal.Type) TypeFood
}

func (f *Feeder) FoodType(animalType animal.Type) TypeFood {
	switch animalType {
	case animal.Cow:
		return Grass
	case animal.Bear:
		return Berries
	case animal.Tiger:
		return Meat
	default:
		return GeneralFood
	}
}

// how much food an animal needs
func (f *Feeder) calculateFood(a animal.Animal) int {
	amount := 0
	switch a.Type {
	case animal.Cow:
		amount = 3
	case animal.Bear:
		amount = 5
	case animal.Tiger:
		amount = 4
	default:
		amount = 2
	}

	return amount * a.Count
}

func (f *Feeder) distributionFood(typeFood TypeFood, amount int) {
	fmt.Printf("%s feed was given in the amount of %d\n", typeFood, amount)
}

func (f *Feeder) FeedAnimals(animals []animal.Animal) {
	for _, beast := range animals {
		foodType := f.FoodType(beast.Type)
		foodNeeded := f.calculateFood(beast)

		if foodAvailable, ok := f.FoodStock[foodType]; ok && foodAvailable >= foodNeeded {
			f.FoodStock[foodType] -= foodNeeded
			f.distributionFood(foodType, foodNeeded)
		} else {
			fmt.Printf("Not enough feed for %s\n", beast.Type)
		}
	}
}
