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

type ZooKeeper struct {
	Name           string
	NumberOfFound  int
	NumberOfCaught int
}

func (k *ZooKeeper) Begin(catch Catcher, cage *Cage) {
	if catch.FindEscapedAnimal() {
		k.NumberOfFound++
		if catch.AttemptCatchAnimal(cage) {
			k.NumberOfCaught++
			cage.CurrentSeats++
		}
	}
}

type Catcher interface {
	FindEscapedAnimal() bool       // Спроба знайти звіра
	AttemptCatchAnimal(*Cage) bool // Сброба зловити звіра в клітку
}

type Animal struct {
	ID     int
	Name   string
	Gender string
	Weight int
	OurNil *int // Вимога задачі використати nil
	Cage        // Так ми розуміємо чи спіймали звіра
}

type Cage struct {
	MaxSeats      int
	MaxWeight     int
	CurrentSeats  int
	CurrentWeight int
}

// Шукаємо звіра. Результат = Знаходимо або ні
func (a *Animal) FindEscapedAnimal() bool {
	return rand.N(2) == 1
}

// Намагаемося зловити в клітку знайденого звіра, якщо в клітці є місця  і не перевищує максимально дозволену вагу
func (a *Animal) AttemptCatchAnimal(cage *Cage) bool {
	// TODO: перевірити чи  правильні тут розрахунки
	if cage.MaxWeight >= a.Weight+cage.CurrentWeight && cage.MaxSeats >= cage.CurrentSeats+1 {
		cage.CurrentWeight += a.Weight
		a.Cage = *cage
		return true
	}
	return false
}

// Функція розможноження
func Reproduction(male, famale *Animal) {
	fakeData := NewFake()
	avgWeight := (male.Weight + famale.Weight) / 2
	NewAnimals := &Animal{ID: escapedAnimalsCount + 1, Name: fakeData.Name, Gender: fakeData.Gender, Weight: avgWeight, Cage: Cage{male.MaxSeats, male.MaxWeight, male.CurrentSeats + 1, male.CurrentWeight + avgWeight}}

	_, err := fmt.Fprintf(os.Stdout, "Replay between %s and %s is complete!! Congratulations to the new animal %v \n", male.Name, famale.Name, *NewAnimals)
	if err != nil {
		return
	}
}

func RandInt(minValue, maxValue int) int {
	return rand.IntN(maxValue-minValue) + minValue
}

// New - функція для створення нових початкових обʼєктів
func New(zkName string, animNumber int) (*ZooKeeper, *Cage) {
	numberSeats := RandInt(1, animNumber+1)    // Дамо звірям шанс
	maxWeight := numberSeats * RandInt(10, 50) // Максимальная вага одного звіра в зоопарку 50
	return &ZooKeeper{zkName, 0, 0}, &Cage{numberSeats, maxWeight, 0, 0}
}

func main() {
	keeper, cage := New("John Wick", escapedAnimalsCount)

	//Нічого іншого не вигадав. Перевести структуру Animal в "type Animal []struct" ломає функцію keeper.Begin -  не  зміг вирулити(  Буду вдячний за підказку!
	var cagedAnimals []Animal

	for i := 0; i < escapedAnimalsCount; i++ {
		fakeData := NewFake()
		animals := &Animal{ID: i + 1, Name: fakeData.Name, Gender: fakeData.Gender, Weight: RandInt(10, 50)}
		keeper.Begin(animals, cage)
		if animals.Cage.MaxSeats > 0 {
			cagedAnimals = append(cagedAnimals, *animals)
		}
	}

	_, err := fmt.Fprintf(os.Stdout, "Zookeeper found %d and cought %d animals Out of %d\n", keeper.NumberOfFound, keeper.NumberOfCaught, escapedAnimalsCount)
	if err != nil {
		return
	}

	//for reproduction:
	var male, famale = -1, -1
	for i, v := range cagedAnimals {
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
		Reproduction(&cagedAnimals[male], &cagedAnimals[famale])
	} else {
		_, err := fmt.Fprintf(os.Stdout, "same-sex reproduction is prohibited")
		if err != nil {
			return
		}
	}
}
