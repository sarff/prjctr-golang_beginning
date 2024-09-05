package animal

type Animal interface {
	Typify() string
	RecordBreathing(breathingData float64)
	RecordSound(soundData float64)
}

type AnimalData struct {
	Type        string
	Pulse       int
	Temperature float64
	Breathing   []float64
	Sounding    []float64
}

func (a *AnimalData) Typify() string {
	return a.Type
}

func (a *AnimalData) RecordBreathing(breathingData float64) {
	a.Breathing = append(a.Breathing, breathingData)
}

func (a *AnimalData) RecordSound(soundData float64) {
	a.Sounding = append(a.Sounding, soundData)
}
