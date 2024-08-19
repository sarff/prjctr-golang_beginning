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
		"hoofed", map[string]Sector{
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
		"birds", map[string]Sector{
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
		"tech room", map[string]Sector{
			"Dining room": {
				ID: 5,
				Animals: []Animal{
					{ID: 13, Name: "Hawk4"},
					//{ID: 2, Name: "Dove2"},
					//{ID: 3, Name: "Dove3"},
				},
			},
			"Restroom": {
				ID: 6,
				Animals: []Animal{
					{ID: 14, Name: "Cow4"},
					//{ID: 2, Name: "Hawk2"},
					//{ID: 3, Name: "Hawk3"},
				},
			},
		},
	},
}

func (z *Zoo) Lookup(name string) (*Sector, *Animal) {
	for _, area := range z.Areas {
		for _, sector := range area.Sectors {
			for _, animal := range sector.Animals {
				if animal.Name == name {
					return &sector, &animal
				}
			}
		}
	}
	return nil, nil
}

func (z *Zoo) Migration(toArea string, toSector int, animal Animal) error {
	for _, area := range z.Areas {
		if area.Name == toArea {
			for sectorName, sector := range area.Sectors {
				if sector.ID == toSector {
					sector.Animals = append(sector.Animals, animal)
					area.Sectors[sectorName] = sector // перезаписуємо сектор
					return nil
				}
			}
		}
	}
	return fmt.Errorf("sector not found")
}

func Cleaning() {
}

func Feeding() {
}

func (z *Zoo) delAnimal(ID int) {
	for _, area := range z.Areas {
		for sectorName, sector := range area.Sectors {
			for i, animal := range sector.Animals {
				if ID == animal.ID {
					fmt.Println("time found and successfully deleted: ", animal)
					sector.Animals = append(sector.Animals[:i], sector.Animals[i+1:]...)
					area.Sectors[sectorName] = sector //перезаписуємо сектор. Без цього дублює останній елемент
					break
				}
			}
		}
	}
}

func main() {
	zoo := NewZoo()

	sector, animal := zoo.Lookup("Hawk3")
	if sector != nil {
		fmt.Printf("%s found in sector.ID = %v\n", animal.Name, sector.ID)
	}

	err := zoo.Migration("tech room", 5, *animal)
	if err != nil {
		fmt.Println(err)
	}
	zoo.delAnimal(10)
	fmt.Println(zoo.Areas)

	//if val, ok := zoo.Areas["Cows"]; ok {
	//	fmt.Println(val)
	//}
}
