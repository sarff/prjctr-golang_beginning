package animal

type AnimalType string

const (
	Bear  AnimalType = "bear"
	Tiger AnimalType = "tiger"
	Lion  AnimalType = "lion"
)

//type Animal interface {
//	Typify() string
//	RecordBreathing(breathingData float64)
//	RecordSound(soundData float64)
//}

type DataAnimal[T float64] struct {
	AnimalType  string
	Pulse       int
	Temperature float64
	Breaths     []T
	Sounds      []T
}

func New(animalPulse int, animalTemperature float64, animalType AnimalType) DataAnimal[float64] {
	return DataAnimal[float64]{
		AnimalType:  string(animalType),
		Pulse:       animalPulse,
		Temperature: animalTemperature,
		Breaths:     []float64{},
		Sounds:      []float64{},
	}
}

func (a *DataAnimal[T]) Typify() string {
	// simulation of animal detection
	if a.Breaths != nil && a.Sounds != nil {
		return a.AnimalType
	}
	return "unknown animal type"
}

func (a *DataAnimal[T]) RecordBreathing(breathingData T) {
	a.Breaths = append(a.Breaths, breathingData)
}

func (a *DataAnimal[T]) RecordSound(soundData T) {
	a.Sounds = append(a.Sounds, soundData)
}
