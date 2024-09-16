package feeder

type Feeder struct {
	DistributionAnimalsInZone DistributionAnimalsInZone
}

type DistributionAnimalsInZone interface {
	AnimalsCalculate()
}

func (f *Feeder) AnimalsCalculate() {
}
