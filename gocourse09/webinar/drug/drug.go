package drug

import "fmt"

type ControllerDrug interface {
	AddDrug(drugName string, dose float64)
	MixDrug(drugNames []string, doses []float64)
}

type Drug struct {
	ControllerDrug ControllerDrug
}

func (d *Drug) AddDrug(drugName string, dose float64) {
	fmt.Printf("Added %s - %f\n", drugName, dose)
}

func (d *Drug) MixDrug(drugNames []string, doses []float64) {
	for i, drug := range drugNames {
		fmt.Printf("Mix %s - %f\n", drug, doses[i])
	}
}
