package collar

import (
	"fmt"
	"math/rand/v2"
	"sync"

	"github.com/sarff/prjctr-golang_beginning/gocourse08/collartask/animal"
)

type SensorAnimal interface {
	Typify() string
	RecordBreathing(breathingData float64)
	RecordSound(soundData float64)
}

type Collar[T float64 | float32] struct {
	mu           sync.RWMutex
	SaveStrategy SaveStrategy[T]
	dataChan     chan animal.DataAnimal[T]
}

type SaveStrategy[T float64 | float32] interface {
	CheckGprs(dataAnimal animal.DataAnimal[T]) string
}

type GprsOff[T float64 | float32] struct{}

func (g *GprsOff[T]) CheckGprs(dataAnimal animal.DataAnimal[T]) string {
	return fmt.Sprintf("%s:Saved to local storage: \nBreathing --> %v \nSounding --> %v ", dataAnimal.AnimalType, dataAnimal.Breaths, dataAnimal.Sounds)
}

type GprsOn[T float64 | float32] struct{}

func (g *GprsOn[T]) CheckGprs(dataAnimal animal.DataAnimal[T]) string {
	return fmt.Sprintf("%s: Saved to server storage: \nBreathing --> %v \nSounding --> %v ", dataAnimal.AnimalType, dataAnimal.Breaths, dataAnimal.Sounds)
}

func (c *Collar[T]) SetStrategy(strategy SaveStrategy[T]) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.SaveStrategy = strategy
}

func (c *Collar[T]) Save(a animal.DataAnimal[T]) {
	c.dataChan <- a
}

func (c *Collar[T]) ProcessData() {
	go func() {
		for data := range c.dataChan {
			c.mu.Lock()
			result := c.SaveStrategy.CheckGprs(data)
			c.mu.Unlock()
			fmt.Println(result)
		}
	}()
}

func (c *Collar[T]) WearCollar(animalData animal.DataAnimal[T]) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			animalData.RecordBreathing(T(rand.Float64() * 100))
			animalData.RecordSound(T(rand.Float64() * 100))

			// Встановлюємо стратегію збереження
			c.SetStrategy(&GprsOn[T]{})
			if rand.IntN(2) == 1 {
				c.SetStrategy(&GprsOff[T]{})
			}

			c.dataChan <- animalData
		}
	}()

	c.ProcessData()
	wg.Wait()
}

func NewCollar[T float64 | float32]() *Collar[T] {
	return &Collar[T]{
		dataChan: make(chan animal.DataAnimal[T]),
	}
}
