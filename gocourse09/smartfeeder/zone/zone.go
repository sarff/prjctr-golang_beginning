package zone

import "github.com/sarff/prjctr-golang_beginning/gocourse09/smartfeeder/animal"

type Zone struct {
	AnimalsInZone []animal.Animal
}

type AnimalsInZone interface {
	CheckZone()
}

func (z Zone) AddAnimals() *Zone {
	animalsInZone := []animal.Animal{
		{Type: animal.Tiger, Count: 3},
		{Type: animal.Cow, Count: 2},
		{Type: animal.Bear, Count: 2},
		{Type: animal.Panda, Count: 22},
	}
	z.AnimalsInZone = append(z.AnimalsInZone, animalsInZone...)
	return &z
}
