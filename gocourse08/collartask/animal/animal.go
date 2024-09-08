package animal

import "math/rand/v2"

type AnimalType string

const (
	Bear  AnimalType = "bear"
	Tiger AnimalType = "tiger"
	Lion  AnimalType = "lion"
)

type Animal interface {
	Typify() string
	RecordBreathing(breathingData float64)
	RecordSound(soundData float64)
}

type DataAnimal[T float64] struct {
	AnimalType  string
	Pulse       int
	Temperature float64
	Breathing   []T
	Sounding    []T
}

func NewAnimal() DataAnimal[float64] {
	animals := [...]AnimalType{Bear, Tiger, Lion}
	return DataAnimal[float64]{
		AnimalType:  string(animals[rand.IntN(len(animals))]),
		Pulse:       rand.IntN(100),
		Temperature: rand.Float64(),
		Breathing:   []float64{},
		Sounding:    []float64{},
	}
}

func (a *DataAnimal[T]) Typify() string {
	// simulation of animal detection
	if a.Breathing != nil && a.Sounding != nil {
		return a.AnimalType
	}
	return "unknown animal type"
}

func (a *DataAnimal[T]) RecordBreathing(breathingData T) {
	a.Breathing = append(a.Breathing, breathingData)
}

func (a *DataAnimal[T]) RecordSound(soundData T) {
	a.Sounding = append(a.Sounding, soundData)
}
