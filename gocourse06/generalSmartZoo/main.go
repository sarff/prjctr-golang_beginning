/*
«Загальна система розумного зоопарку»
Концепція
Створити програму для управління розумним зоопарком, де декілька горутин виконують різні завдання, такі як моніторинг
стану тварин, керування доступом до вольєрів та управління кормушками. Програма має активно використовувати канали для
комунікації між горутинами, уникаючи «race conditions» і «deadlocks».
Завдання
Моніторинг стану тварин: Створіть горутину для кожної тварини в зоопарку. Кожна горутина збирає дані про стан тварини
(наприклад, рівень здоров'я, голод, настрій) і відправляє їх через канал до центральної системи моніторингу.
Керування доступом до вольєрів: Імплементуйте горутину, яка контролює доступ до вольєрів, використовуючи канали для
отримання запитів на відкриття/закриття.
Управління кормушками: Розробіть горутини для управління автоматичними кормушками, які відправляють статус кормушки
(порожня/повна) через канал.
Умови виконання
Уникнення «Race Conditions»: забезпечте, щоб спільні ресурси (наприклад, дані про стан тварин) були захищені від
одночасного доступу декількома горутинами. Використовуйте канали для синхронізації доступу.
Управління «Deadlocks»: уважно використовуйте блокування та канали, щоб уникнути взаємних блокувань між горутинами.
Логування та моніторинг: реалізуйте систему логування, яка фіксує важливі події у системі, наприклад, коли тварина
потребує уваги або коли кормушка порожня.
Тестування: напишіть модульні тести для перевірки коректності взаємодії між горутинами та уникнення «race conditions»
та «deadlocks».
*/
package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Animal описує структуру даних для тварини
type Animal struct {
	ID     int
	Health int
	Hunger int
	Mood   int
}

// Enclosure описує статус вольєра
type Enclosure struct {
	ID     int
	IsOpen bool
}

// Feeder описує статус кормушки
type Feeder struct {
	ID      int
	IsEmpty bool
}

// Генерує тестові дані для тварин
func generateAnimals(n int) []Animal {
	var animals []Animal
	for i := 0; i < n; i++ {
		animal := Animal{
			ID:     i,
			Health: rand.IntN(100),
			Hunger: rand.IntN(100),
			Mood:   rand.IntN(100),
		}
		animals = append(animals, animal)
	}
	return animals
}

// Генерує тестові дані для вольєрів
func generateEnclosures(n int) []Enclosure {
	var enclosures []Enclosure
	for i := 0; i < n; i++ {
		enclosure := Enclosure{
			ID:     i,
			IsOpen: rand.IntN(2) == 1,
		}
		enclosures = append(enclosures, enclosure)
	}
	return enclosures
}

// Генерує тестові дані для кормушок
func generateFeeders(n int) []Feeder {
	var feeders []Feeder
	for i := 0; i < n; i++ {
		feeder := Feeder{
			ID:      i,
			IsEmpty: rand.IntN(2) == 1,
		}
		feeders = append(feeders, feeder)
	}
	return feeders
}

func controlCondition(animalChan chan *Animal, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)

	select {
	case animal, ok := <-animalChan:
		if ok {
			if animal.Health < 50 {
				fmt.Printf("Animal with ID %v - needs help\n", animal.ID)
			}
			if animal.Hunger < 30 {
				fmt.Printf("Animal with ID %v - needs to be fed\n", animal.ID)
			}
			if animal.Mood < 30 {
				fmt.Printf("Animal with ID %v - needs to be released from the cage\n", animal.ID)
			}
		}
	}

}

func controlEnclosure(enclosures Enclosure) {
	_ = enclosures.IsOpen
}

func controlFeeder(feeders Feeder) {
	_ = feeders.IsEmpty
}

func main() {
	wg := new(sync.WaitGroup)

	// Генеруємо тестові дані
	animals := generateAnimals(10)
	enclosures := generateEnclosures(5)
	feeders := generateFeeders(3)

	// Виводимо згенеровані дані
	fmt.Println("Animals:", animals)
	fmt.Println("Enclosures:", enclosures)
	fmt.Println("Feeders:", feeders)

	animalChan := make(chan *Animal)

	for _, animal := range animals {
		wg.Add(1)
		go controlCondition(animalChan, wg)
		animalChan <- &animal
	}
	wg.Wait()

	for _, enclosure := range enclosures {
		go controlEnclosure(enclosure)
	}

	for _, feeder := range feeders {
		go controlFeeder(feeder)
	}

}
