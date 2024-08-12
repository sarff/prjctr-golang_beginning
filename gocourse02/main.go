package main

import (
	"fmt"
	"math/rand/v2"
	"os"
)

/*
Написати програму “Зоопарк”. Звіри повтікали (більше трьох штук), наглядач повинен їх зібрати.
Кожна сутність (наглядач, звір, клітка, …) представляється окремою структурою (zookeeper, animal, cage, …).
Треба використати щонайменше: структури, вказівник, nil, будування, конструктор. Додати тваринам можливість
розмножуватись. Програма має демонструвати свою роботу через вивід в stdout.
*/

const escapedAnimalsCount = 15 // Кількість звірів які повтікали
//type Gender string
//
//const (
//	male   Gender = "male"
//	female Gender = "female"
//)

type ZooKeeper struct {
	Name          string
	AnimalsFound  int
	AnimalsCaught int
}

func (k *ZooKeeper) SearchingForAnimalsReturnToCage(catch Catcher, animal *Animal, cage *Cage) {
	if catch.FindEscapedAnimal() {
		k.AnimalsFound++
		if catch.AttemptCatchAnimal(cage, animal) {
			k.AnimalsCaught++
			cage.CurrentSeats++
		}
	}
}

type Catcher interface {
	FindEscapedAnimal() bool                // Спроба знайти звіра
	AttemptCatchAnimal(*Cage, *Animal) bool // Сброба зловити звіра в клітку
}

type Animal struct {
	ID     int
	Name   string
	Gender string
	Weight int
	OurNil *int // Вимога задачі використати nil
	//Cage        // Так ми розуміємо чи спіймали звіра
}

type Cage struct {
	MaxSeats      int
	MaxWeight     int
	CurrentSeats  int
	CurrentWeight int
	Animal        []Animal
}

// Шукаємо звіра. Результат = Знаходимо або ні
func (k *ZooKeeper) FindEscapedAnimal() bool {
	return rand.N(2) == 1
}

// Намагаемося зловити в клітку знайденого звіра, якщо в клітці є місця  і не перевищує максимально дозволену вагу
func (k *ZooKeeper) AttemptCatchAnimal(cage *Cage, a *Animal) bool {
	if cage.MaxWeight >= a.Weight+cage.CurrentWeight && cage.MaxSeats >= cage.CurrentSeats+1 {
		cage.CurrentWeight += a.Weight
		cage.Animal = append(cage.Animal, *a)
		return true
	}
	return false
}

// Функція розможноження
func (male *Animal) Reproduction(famale *Animal) {
	fakeData := NewFake()
	avgWeight := (male.Weight + famale.Weight) / 2
	NewAnimals := &Animal{ID: escapedAnimalsCount + 1, Name: fakeData.Name, Gender: fakeData.Gender, Weight: avgWeight}

	fmt.Fprintf(os.Stdout, "Replay between %s and %s is complete!! Congratulations to the new animal %v \n", male.Name, famale.Name, *NewAnimals)

}

func RandInt(minValue, maxValue int) int {
	return rand.IntN(maxValue-minValue) + minValue
}

func NewZooKeeper(zkName string) *ZooKeeper {
	return &ZooKeeper{zkName, 0, 0}
}

func NewCage(animNumber int) *Cage {
	numberSeats := RandInt(1, animNumber+1)    // Дамо звірям шанс
	maxWeight := numberSeats * RandInt(10, 50) // Максимальная вага одного звіра в зоопарку 50
	animals := make([]Animal, 0)
	return &Cage{numberSeats, maxWeight, 0, 0, animals}
}

func NewAnimal(i int) *Animal {
	fakeData := NewFake()
	return &Animal{ID: i + 1, Name: fakeData.Name, Gender: fakeData.Gender, Weight: RandInt(10, 50)}
}

func main() {
	keeper := NewZooKeeper("John Wick")
	cage := NewCage(escapedAnimalsCount)

	for i := 0; i < escapedAnimalsCount; i++ {
		animal := NewAnimal(i)
		keeper.SearchingForAnimalsReturnToCage(keeper, animal, cage)
	}

	fmt.Fprintf(os.Stdout, "Zookeeper found %d and caught %d animals Out of %d\n", keeper.AnimalsFound, keeper.AnimalsCaught, escapedAnimalsCount)

	// for reproduction:
	var male, famale = -1, -1
	for i, v := range cage.Animal {
		switch v.Gender {
		case "male":
			if male == -1 {
				male = i
			}
		case "female":
			if famale == -1 {
				famale = i
			}
		}
		fmt.Println(i, v)
	}

	if male > -1 && famale > -1 {
		cage.Animal[male].Reproduction(&cage.Animal[famale])
	} else {
		fmt.Println("same-sex reproduction is prohibited")
	}
}
