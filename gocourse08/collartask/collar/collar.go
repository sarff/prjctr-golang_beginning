package collar

import (
	"fmt"
	"sync"

	"github.com/sarff/prjctr-golang_beginning/gocourse08/collartask/animal"
)

type Collar struct {
	SaveStrategy SaveStrategy
	dataChan     chan animal.DataAnimal[float64]
}

type SaveStrategy interface {
	CheckGprs(dataAnimal animal.DataAnimal[float64]) string
}

type GprsOff struct{}

func (g *GprsOff) CheckGprs(dataAnimal animal.DataAnimal[float64]) string {
	return fmt.Sprintf("Saved to local storage: \nBreathing --> %v \nSounding --> %v ", dataAnimal.Breathing, dataAnimal.Sounding)
}

type GprsOn struct{}

func (g *GprsOn) CheckGprs(dataAnimal animal.DataAnimal[float64]) string {
	return fmt.Sprintf("Saved to server storage: \nBreathing --> %v \nSounding --> %v ", dataAnimal.Breathing, dataAnimal.Sounding)
}

func (c *Collar) SetStrategy(strategy SaveStrategy) {
	c.SaveStrategy = strategy
}

func (c *Collar) Save(a animal.DataAnimal[float64]) {
	c.dataChan <- a
}

func (c *Collar) ProcessData() {
	mu := sync.Mutex{}
	for data := range c.dataChan {
		mu.Lock()
		result := c.SaveStrategy.CheckGprs(data)
		mu.Unlock()
		fmt.Println(result)
	}
}

func NewCollar() *Collar {
	return &Collar{
		dataChan: make(chan animal.DataAnimal[float64]),
	}
}
