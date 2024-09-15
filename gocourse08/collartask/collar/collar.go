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

type Collar struct {
	SaveStrategy SaveStrategy
	dataChan     chan animal.DataAnimal[float64]
	mu           sync.RWMutex
}

type SaveStrategy interface {
	CheckGprs(dataAnimal animal.DataAnimal[float64]) string
}

type GprsOff struct{}

func (g *GprsOff) CheckGprs(dataAnimal animal.DataAnimal[float64]) string {
	return fmt.Sprintf("%s:Saved to local storage: \nBreathing --> %v \nSounding --> %v ", dataAnimal.AnimalType, dataAnimal.Breaths, dataAnimal.Sounds)
}

type GprsOn struct{}

func (g *GprsOn) CheckGprs(dataAnimal animal.DataAnimal[float64]) string {
	return fmt.Sprintf("%s: Saved to server storage: \nBreathing --> %v \nSounding --> %v ", dataAnimal.AnimalType, dataAnimal.Breaths, dataAnimal.Sounds)
}

func (c *Collar) SetStrategy(strategy SaveStrategy) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.SaveStrategy = strategy
}

func (c *Collar) Save(a *animal.DataAnimal[float64]) {
	c.dataChan <- *a
}

func (c *Collar) ProcessData() {
	go func() {
		for data := range c.dataChan {
			c.mu.Lock()
			result := c.SaveStrategy.CheckGprs(data)
			c.mu.Unlock()
			fmt.Println(result)
		}
	}()
}

func (c *Collar) WearCollar(animalData animal.DataAnimal[float64]) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			animalData.RecordBreathing(rand.Float64() * 100)
			animalData.RecordSound(rand.Float64() * 100)

			// Встановлюємо стратегію збереження
			c.SetStrategy(&GprsOn{})
			if rand.IntN(2) == 1 {
				c.SetStrategy(&GprsOff{})
			}

			c.dataChan <- animalData
		}
	}()

	c.ProcessData()
	wg.Wait()
}

func NewCollar() *Collar {
	return &Collar{
		dataChan: make(chan animal.DataAnimal[float64]),
	}
}
