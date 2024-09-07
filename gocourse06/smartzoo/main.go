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
	"log/slog"
	"math/rand/v2"
	"os"
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
	animals := make([]Animal, n)
	for i := range n {
		animals[i] = Animal{
			ID:     i,
			Health: rand.IntN(100),
			Hunger: rand.IntN(100),
			Mood:   rand.IntN(100),
		}
	}
	return animals
}

// Генерує тестові дані для вольєрів
func generateEnclosures(n int) []Enclosure {
	enclosures := make([]Enclosure, n)
	for i := range n {
		enclosures[i] = Enclosure{
			ID:     i,
			IsOpen: rand.IntN(2) == 1,
		}
	}
	return enclosures
}

// Генерує тестові дані для годівниці
func generateFeeders(n int) []Feeder {
	feeders := make([]Feeder, n)
	for i := range n {
		feeders[i] = Feeder{
			ID:      i,
			IsEmpty: rand.IntN(2) == 1,
		}
	}
	return feeders
}

func controlCondition(animalChan <-chan *Animal, wg *sync.WaitGroup, log *slog.Logger) {
	defer wg.Done()
	wg.Add(1)
	time.Sleep(1 * time.Second)

	if animal, ok := <-animalChan; ok {
		if animal.Health < 30 {
			log.Warn("Animal needs help", slog.Int("animalId", animal.ID))
		}
		if animal.Hunger < 30 {
			log.Warn("Animal needs to be fed", slog.Int("animalId", animal.ID))
		}
		if animal.Mood < 30 {
			log.Warn("Animal - needs to be released from the cage", slog.Int("animalId", animal.ID))
		}
	}
}

func controlEnclosure(enclosureChan chan *Enclosure, isOpen bool, wg *sync.WaitGroup, log *slog.Logger) {
	defer wg.Done()
	wg.Add(1)
	if enclosure, ok := <-enclosureChan; ok {
		enclosure.IsOpen = isOpen
		if isOpen {
			log.Info("Closing the enclosure...")
		} else {
			log.Info("Opening the enclosure...")
		}
		time.Sleep(1 * time.Second)
	}
}

func controlFeeder(feedChan chan *Feeder, wg *sync.WaitGroup, log *slog.Logger) {
	defer wg.Done()
	wg.Add(1)
	time.Sleep(3 * time.Second) // must be more than 3
	if feed, ok := <-feedChan; ok {
		if feed.IsEmpty {
			log.Warn("Needs a refill feeder", slog.Int("feedId", feed.ID))
		}
	}
}

func loggerNew(writer *os.File) *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(writer, nil))
	slog.SetDefault(logger)
	return logger
}

func main() {
	log := loggerNew(os.Stdout)
	wg := new(sync.WaitGroup)

	// Генеруємо тестові дані
	animals := generateAnimals(5)
	enclosures := generateEnclosures(5)
	feeders := generateFeeders(5)

	// Виводимо згенеровані дані
	fmt.Println("Animals:", animals)
	fmt.Println("Enclosures:", enclosures)
	fmt.Println("Feeders:", feeders)

	animalChan := make(chan *Animal)
	enclosureChan := make(chan *Enclosure)
	feederChan := make(chan *Feeder)
	for _, animal := range animals {
		go controlCondition(animalChan, wg, log)
		animalChan <- &animal
	}

	for _, enclosure := range enclosures {
		isOpen := true
		if enclosure.IsOpen {
			isOpen = false
		}
		go controlEnclosure(enclosureChan, isOpen, wg, log)
		enclosureChan <- &enclosure
	}

	for _, feeder := range feeders {
		go controlFeeder(feederChan, wg, log)
		feederChan <- &feeder
	}

	log.Info("Simulation Done")
}
