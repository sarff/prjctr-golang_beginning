package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/sarff/prjctr-golang_beginning/gocourse09/smartfeeder/animal"
	"github.com/sarff/prjctr-golang_beginning/gocourse09/smartfeeder/feeder"
	"github.com/sarff/prjctr-golang_beginning/gocourse09/smartfeeder/zone"
)

/*
На великій земельній ділянці заповідника стоїть велике сховище різних кормів для копитних тварин.
До нього підʼєднана електрона кормушка, яка аналізує стан біля себе: є тварини чи немає, які саме і скільки, якщо є.
У залежності від стану видає потрібну кількість потрібної їжі.
Варіанти сутностей:
Зона знаходження тварин, які харчуються: якщо тварина в цій зоні — кормушка видає корм.
Аналізатор зони знаходження: створює перелік тварин у зоні в момент часу.
Тварина: одиниця тварини.
Брекет корму: одиниця корму для певної тварини.
Притримуватися SOLID. Написати модульні тести. Можна використовувати горутини та генеріки.
*/

func main() {
	// initial state of storage
	foodStock := map[feeder.TypeFood]int{
		feeder.Berries:     rand.IntN(20),
		feeder.Grass:       rand.IntN(25),
		feeder.Meat:        rand.IntN(15),
		feeder.GeneralFooD: rand.IntN(45),
	}

	feederMain := &feeder.Feeder{
		FoodStock: foodStock,
	}

	zoneSetUp := zone.Zone{AnimalsInZone: make([]animal.Animal, 0)}.CheckZone()

	fmt.Printf("The following animals are near the feeder: %v\n", zoneSetUp.AnimalsInZone)

	feederMain.FeedAnimals(zoneSetUp.AnimalsInZone)
}
