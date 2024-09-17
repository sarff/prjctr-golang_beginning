package feeder

import (
	"testing"

	"github.com/sarff/prjctr-golang_beginning/gocourse09/smartfeeder/animal"
)

func TestFeeder_FoodType(t *testing.T) {
	foodStock := map[TypeFood]int{
		Berries:     20,
		Grass:       25,
		Meat:        15,
		GeneralFooD: 45,
	}

	feederTest := &Feeder{
		FoodStock: foodStock,
	}

	valid := map[animal.Type]TypeFood{
		animal.Panda: GeneralFooD,
		animal.Cow:   Grass,
		animal.Bear:  Berries,
		animal.Tiger: Meat,
	}

	for animalType, typeFood := range valid {
		t.Run(string(typeFood), func(t *testing.T) {
			getFood := feederTest.FoodType(animalType)
			if getFood != typeFood {
				t.Errorf("FoodType should have returned %s, but return %s", typeFood, getFood)
			}
		})
	}
}
