package animal

type AnimalType string

const (
	Bear    AnimalType = "bear"
	Tiger   AnimalType = "tiger"
	Lion    AnimalType = "lion"
	Unknown AnimalType = "unknown"
)

type DataAnimal[T float64 | float32] struct {
	AnimalType  AnimalType
	Pulse       int
	Temperature float64
	Breaths     []T
	Sounds      []T
}

func New[T float64 | float32](pulse int, temperature float64, animalType AnimalType) DataAnimal[T] {
	return DataAnimal[T]{
		AnimalType:  animalType,
		Pulse:       pulse,
		Temperature: temperature,
		Breaths:     []T{},
		Sounds:      []T{},
	}
}

func (a *DataAnimal[T]) Typify() AnimalType {
	// simulation of animal detection
	if a.Breaths != nil && a.Sounds != nil {
		return a.AnimalType
	}
	return Unknown
}

func (a *DataAnimal[T]) RecordBreathing(breathingData T) {
	a.Breaths = append(a.Breaths, breathingData)
}

func (a *DataAnimal[T]) RecordSound(soundData T) {
	a.Sounds = append(a.Sounds, soundData)
}
