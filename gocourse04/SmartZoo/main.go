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

func NewZoo() *Zoo {
	return &Zoo{Areas: AreasData}
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

// Returns the found animal and in which sector and area it is
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
func (z *Zoo) Migration(fromAreaName, toArea string, fromSectorID, toSector int, animal Animal) error {
	for _, area := range z.Areas {
		if area.Name == toArea {
			for sectorName, sector := range area.Sectors {
				if sector.ID == toSector {
					sector.Animals = append(sector.Animals, animal)
					area.Sectors[sectorName] = sector // rewrite sector
					z.delOldRecord(fromAreaName, fromSectorID, animal)
					return nil
				}
			}
		}
	}
	return fmt.Errorf("sector not found")
}

func (z *Zoo) delOldRecord(fromAreaName string, fromSectorID int, animal Animal) {
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
func (a Animal) Cleaning(fromArea string, sectorID int, z Zoo) {
	var toArea string
	var toSector int
	areaClean := "tech room"
	sectorClean := 6
	for _, area := range z.Areas {
		if area.Name == areaClean {
			for sectorName, sector := range area.Sectors {
				if sector.ID == sectorClean {
					toSector = sector.ID
					toArea = area.Name
					sector.Animals = append(sector.Animals, a)
					area.Sectors[sectorName] = sector // rewrite sector
					z.delOldRecord(fromArea, sectorID, a)
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
					z.delOldRecord(toArea, toSector, a)
				}
			}
		}
	}
}

// the animals go out to eat - Sector: Dining room
func (z *Zoo) Feeding(fromaArea, fromSector string) error {
	diningRoom := z.Areas["Techroom"].Sectors["Dining room"]

	if val, ok := z.Areas[fromaArea].Sectors[fromSector]; ok {
		for _, sector := range val.Animals {
			diningRoom.Animals = append(diningRoom.Animals, sector)
		}
		z.Areas["Techroom"].Sectors["Dining room"] = diningRoom // rewrite sector
		z.Areas[fromaArea].Sectors[fromSector] = Sector{}       // emptying the sector from which all the animals have left
	} else {
		return errors.New("sector in this Area - not found")
	}
	return nil
}

func (z *Zoo) delAnimal(ID int) string {
	for _, area := range z.Areas {
		for sectorName, sector := range area.Sectors {
			for i, animal := range sector.Animals {
				if ID == animal.ID {
					sector.Animals = append(sector.Animals[:i], sector.Animals[i+1:]...)
					area.Sectors[sectorName] = sector // rewrite sector
					return fmt.Sprintf("time found and successfully deleted: %s", animal.Name)
				}
			}
		}
	}
	return ""
}

func main() {
	zoo := NewZoo()
	zooAreas := zoo.Areas
	animalName := "Hawk3"
	animalArea, animalSector, animal := zoo.Lookup(animalName)
	if animal != nil {
		fmt.Printf("%s found in animalSector.ID = %v\n", animal.Name, animalSector.ID)
	} else {
		fmt.Printf("%s not found\n", animalName)
	}

	if animal != nil {
		fmt.Printf("%s is taken away for hygiene procedures\n", animal.Name)
		animal.Cleaning(animalArea.Name, animalSector.ID, *zoo)
		fmt.Printf("%s has returned to his sector\n", animal.Name)
	}

	if animal != nil {
		err := zoo.Migration(animalArea.Name, "tech room", animalSector.ID, 5, *animal)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("The array after the Migration: ", zooAreas)
	}

	err := zoo.Feeding("Hoofed", "Cows")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The animals are eating")
	}

	fmt.Println(zoo.delAnimal(10))

	fmt.Println("Array after deletion:", zooAreas)
}
