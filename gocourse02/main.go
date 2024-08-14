package main

import (
	"fmt"
	"math/rand/v2"
)

/*
Написати програму “Зоопарк”. Звіри повтікали (більше трьох штук), наглядач повинен їх зібрати.
Кожна сутність (наглядач, звір, клітка, …) представляється окремою структурою (zookeeper, animal, cage, …).
Треба використати щонайменше: структури, вказівник, nil, будування, конструктор. Додати тваринам можливість
розмножуватись. Програма має демонструвати свою роботу через вивід в stdout.
*/

const escapedAnimalsCount = 15 // Кількість звірів які повтікали

type ZooKeeper struct {
	Name          string
	AnimalsFound  int
	AnimalsCaught int
}

type _ interface {
	FindEscapedAnimal() bool                // Спроба знайти звіра
	AttemptCatchAnimal(*Cage, *Animal) bool // Сброба зловити звіра в клітку
}

type Animal struct {
	ID     int
	Name   string
	Gender string
	Weight int
}

type Cage struct {
	MaxSeats      int
	MaxWeight     int
	CurrentSeats  int
	CurrentWeight int
	Animals       []Animal
}

func NewAnimal(id, weight int) *Animal {
	fakeData := NewFake()
	return &Animal{ID: id + 1, Name: fakeData.Name, Gender: fakeData.Gender, Weight: weight}
}

func NewZooKeeper(name string) *ZooKeeper {
	return &ZooKeeper{Name: name, AnimalsFound: 0, AnimalsCaught: 0}
}

func NewCage(capacity int) *Cage {
	numberSeats := RandInt(1, capacity+1)      // Дамо звірям шанс
	maxWeight := numberSeats * RandInt(10, 50) // Максимальная вага одного звіра в зоопарку 50
	// numberSeats+1 - Максимальная кількість звірів яка може поміститися в клітку +1=child
	animals := make([]Animal, 0, numberSeats+1)
	return &Cage{MaxSeats: numberSeats, MaxWeight: maxWeight, CurrentSeats: 0, CurrentWeight: 0, Animals: animals}
}

func (k *ZooKeeper) SearchForAnimalAndReturnToCage(animal *Animal, cage *Cage) {
	if k.FindEscapedAnimal() {
		k.AnimalsFound++
		if k.AttemptCatchAnimal(cage, animal) {
			k.AnimalsCaught++
			cage.CurrentSeats++
		}
	}
}

// Шукаємо звіра. Результат = Знаходимо або ні
func (k *ZooKeeper) FindEscapedAnimal() bool {
	return rand.N(2) == 1
}

// Намагаемося зловити в клітку знайденого звіра, якщо в клітці є місця  і не перевищує максимально дозволену вагу
func (k *ZooKeeper) AttemptCatchAnimal(cage *Cage, animal *Animal) bool {
	if cage.MaxWeight >= animal.Weight+cage.CurrentWeight && cage.MaxSeats >= cage.CurrentSeats+1 {
		cage.CurrentWeight += animal.Weight
		cage.Animals = append(cage.Animals, *animal)
		return true
	}
	return false
}

// Функція розмноження
func (male *Animal) Reproduction(female *Animal) *Animal {
	avgWeight := (male.Weight + female.Weight) / 2
	if rand.N(2) == 1 {
		//child.id завжди буде +1 від константи escapedAnimalsCount, це прописано в функції NewAnimal
		child := NewAnimal(escapedAnimalsCount, avgWeight)
		return child
	} else {
		return nil
	}

}

// RandInt generates integer from <low, high> interval.
func RandInt(low, high int) int {
	return rand.IntN(high-low) + low
}

func main() {
	keeper := NewZooKeeper("John Wick")
	cage := NewCage(escapedAnimalsCount)

	for i := 0; i < escapedAnimalsCount; i++ {
		animal := NewAnimal(i, RandInt(10, 50))
		keeper.SearchForAnimalAndReturnToCage(animal, cage)
	}

	fmt.Printf("Zookeeper found %d and caught %d animals "+
		"Out of %d\n", keeper.AnimalsFound, keeper.AnimalsCaught, escapedAnimalsCount)

	// for reproduction:
	var male, female = -1, -1
	for i, v := range cage.Animals {
		switch v.Gender {
		case "male":
			if male == -1 {
				male = i
			}
		case "female":
			if female == -1 {
				female = i
			}
		}
		fmt.Println(i, v)
	}

	if male > -1 && female > -1 {
		child := cage.Animals[male].Reproduction(&cage.Animals[female])
		if child != nil {
			fmt.Printf("Replay between %s and %s is complete!! "+
				"Congratulations to the new animal %v \n", cage.Animals[male].Name, cage.Animals[female].Name, *child)
		} else {
			fmt.Printf("%s and %s failed to give birth "+
				"to a cub\n", cage.Animals[male].Name, cage.Animals[female].Name)
		}
	} else {
		fmt.Println("same-sex reproduction is prohibited")
	}
}
