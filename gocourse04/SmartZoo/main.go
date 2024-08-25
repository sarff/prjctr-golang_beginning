/*
При розробці системи «Розумний зоопарк» техлід ще не вирішив, яку базу даних використовувати і для реалізації
прототипа вирішили зберігати дані в памʼяті програми, а саме — в мапах. Треба побудувати складну структуру мап,
відобразивши деяку ієрархію частин зоопарку. Наприклад, зоопарк ділиться на території: копитні, пернаті,
примати… та інше. Кожна територія має декілька загонів для кожного виду тварин. Кожен загін може мати ділянку
для перебування тварини і технічне приміщення, де можуть зберігатися приладдя для роботи з тваринами. На кожній
території є певна кількість тварин, а технічне приміщення «здатне» виконувати деякі функції, такі як прибирання і
годування тварин. Реалізувати функції пошуку тварини за імʼям або ID, переміщення тварини із загону в загін,
годування тварини.
*/
package main

import (
	"errors"
	"fmt"
)

const (
	areaClean   = "tech room"
	sectorClean = 6

	toAreaName = "tech room"
	toSectorID = 5

	fromArea   = "Hoofed"
	fromSector = "Cows"

	animalNameLookup = "Hawk3"

	deleteAnimalID = 10

	feedArea   = "Techroom"
	feedSector = "Dining room"
)

type Areas map[string]Area

// Територія по  типу (копитні, пернаті, примати)
type Area struct {
	Name    string
	Sectors map[string]Sector
}

// Загон для виду тварин. У кожному загоні є технічне приміщення (наприклад годування)
type Sector struct {
	ID      int
	Animals []Animal
}

type Animal struct {
	ID   int
	Name string
}

type Zoo struct {
	Areas Areas
}

func NewZoo(areas map[string]Area) *Zoo {
	return &Zoo{Areas: areas}
}

var AreasData = map[string]Area{
	"Hoofed": {
		Name: "hoofed", Sectors: map[string]Sector{
			"Horses": {
				ID: 1,
				Animals: []Animal{
					{ID: 1, Name: "Horse1"},
					{ID: 2, Name: "Horse2"},
					{ID: 3, Name: "Horse3"},
				},
			},
			"Cows": {
				ID: 2,
				Animals: []Animal{
					{ID: 4, Name: "Cow1"},
					{ID: 5, Name: "Cow2"},
					{ID: 6, Name: "Cow3"},
				},
			},
		},
	},
	"Birds": {
		Name: "birds", Sectors: map[string]Sector{
			"Doves": {
				ID: 3,
				Animals: []Animal{
					{ID: 7, Name: "Dove1"},
					{ID: 8, Name: "Dove2"},
					{ID: 9, Name: "Dove3"},
				},
			},
			"Hawks": {
				ID: 4,
				Animals: []Animal{
					{ID: 10, Name: "Hawk1"},
					{ID: 11, Name: "Hawk2"},
					{ID: 12, Name: "Hawk3"},
				},
			},
		},
	},
	"Techroom": {
		Name: "tech room", Sectors: map[string]Sector{
			"Dining room": {
				ID: 5,
				Animals: []Animal{
					{ID: 13, Name: "Hawk4"},
				},
			},
			"Bath room": {
				ID: 6,
				Animals: []Animal{
					{ID: 14, Name: "Cow4"},
				},
			},
		},
	},
}

// Lookup returns the found animal and in which sector and area it is.
func (z *Zoo) Lookup(name string) (*Area, *Sector, *Animal) {
	for _, area := range z.Areas {
		for _, sector := range area.Sectors {
			for _, animal := range sector.Animals {
				if animal.Name == name {
					return &area, &sector, &animal
				}
			}
		}
	}
	return nil, nil, nil
}

// move the animal to another sector
func (z *Zoo) MoveAnimal(fromAreaName, toAreaName string, fromSectorID, toSectorID int, animal Animal) error {
	for _, area := range z.Areas {
		if area.Name == toAreaName {
			for sectorName, sector := range area.Sectors {
				if sector.ID == toSectorID {
					sector.Animals = append(sector.Animals, animal)
					area.Sectors[sectorName] = sector // rewrite sector
					z.deleteOldRecord(fromAreaName, fromSectorID, animal)
					return nil
				}
			}
		}
	}
	return errors.New("sector not found")
}

func (z *Zoo) deleteOldRecord(fromAreaName string, fromSectorID int, animal Animal) {
	for _, area := range z.Areas {
		if area.Name == fromAreaName {
			for sectorName, sector := range area.Sectors {
				if sector.ID == fromSectorID {
					for i, a := range sector.Animals {
						if a.ID == animal.ID {
							sector.Animals = append(sector.Animals[:i], sector.Animals[i+1:]...)
							area.Sectors[sectorName] = sector
						}
					}
				}
			}
		}
	}
}

// clean the animal. ID 6 = Bath room
// the animal moves to the “Bath room” sector and returns to the sector from which it came.
func (a Animal) Clean(fromArea string, sectorID int, z Zoo) {
	var toArea string
	var toSector int
	for _, area := range z.Areas {
		if area.Name == areaClean {
			for sectorName, sector := range area.Sectors {
				if sector.ID == sectorClean {
					toSector = sector.ID
					toArea = area.Name
					sector.Animals = append(sector.Animals, a)
					area.Sectors[sectorName] = sector // rewrite sector
					z.deleteOldRecord(fromArea, sectorID, a)
				}
			}
		}
	}

	// roll back
	for _, area := range z.Areas {
		if area.Name == fromArea {
			for sectorName, sector := range area.Sectors {
				if sector.ID == sectorID {
					sector.Animals = append(sector.Animals, a)
					area.Sectors[sectorName] = sector // rewrite  sector
					z.deleteOldRecord(toArea, toSector, a)
				}
			}
		}
	}
}

// the animals go out to eat - Sector: Dining room
func (z *Zoo) FeedAnimals(fromArea, fromSector string) error {
	diningRoom := z.Areas[feedArea].Sectors[feedSector]
	val, ok := z.Areas[fromArea].Sectors[fromSector]
	if !ok {
		return errors.New("sector in this Area - not found")
	}

	diningRoom.Animals = append(diningRoom.Animals, val.Animals...)
	z.Areas[feedArea].Sectors[feedSector] = diningRoom // rewrite sector
	z.Areas[fromArea].Sectors[fromSector] = Sector{}   // emptying the sector from which all the animals have left

	return nil
}

func (z *Zoo) deleteAnimal(id int) bool {
	for _, area := range z.Areas {
		for sectorName, sector := range area.Sectors {
			for i, animal := range sector.Animals {
				if id == animal.ID {
					sector.Animals = append(sector.Animals[:i], sector.Animals[i+1:]...)
					area.Sectors[sectorName] = sector // rewrite sector
					return true
				}
			}
		}
	}
	return false
}

func main() {
	zoo := NewZoo(AreasData)
	animalArea, animalSector, animal := zoo.Lookup(animalNameLookup)
	if animal != nil {
		fmt.Printf("%s found in animalSector.ID = %v\n", animal.Name, animalSector.ID)

		fmt.Printf("%s is taken away for hygiene procedures\n", animal.Name)
		animal.Clean(animalArea.Name, animalSector.ID, *zoo)
		fmt.Printf("%s has returned to his sector\n", animal.Name)

		err := zoo.MoveAnimal(animalArea.Name, toAreaName, animalSector.ID, toSectorID, *animal)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("The map after the MoveAnimal: ", zoo.Areas)
	} else {
		fmt.Printf("%s not found\n", animalNameLookup)
	}

	err := zoo.FeedAnimals(fromArea, fromSector)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The animals are eating from Area %v and from Sector %v\n", fromArea, fromSector)
	}

	if zoo.deleteAnimal(deleteAnimalID) {
		fmt.Printf("Found and successfully deleted animal with id=%d\n", deleteAnimalID)
	}

	fmt.Println("Map after deletion:", zoo.Areas)
}
