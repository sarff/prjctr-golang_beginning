package dropper

import (
	"fmt"

	"github.com/sarff/prjctr-golang_beginning/gocourse09/webinar/drug"

	"github.com/sarff/prjctr-golang_beginning/gocourse09/webinar/patient"
)

type Dropper struct {
	CheckerPatient patient.CheckerPatient
	ControllerDrug drug.ControllerDrug
}

type Reacter interface {
	React() string
}

func (d *Dropper) React() {
	condition := d.CheckerPatient.CheckCondition()
	switch condition {
	case patient.ConditionGood:
		d.ControllerDrug.AddDrug("Water", 1.5)
	case patient.ConditionNormal:
		d.ControllerDrug.AddDrug("Analgon", 1.5)
	case patient.ConditionCritical:
		d.ControllerDrug.MixDrug([]string{"Drug1", "Drug2"}, []float64{1.01, 2.02})
	default:
		fmt.Println("Unrecognized patientState")
	}
}
