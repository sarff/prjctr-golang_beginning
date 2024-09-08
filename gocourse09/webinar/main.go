/*
Написати програму для електронної крапельниці, яка сама має доступ до кількох джерел ліків і сама їх додає/змішує
взалежності від стану пацієнта.
Притримуванися SOLID
*/
package main

import (
	"github.com/sarff/prjctr-golang_beginning/gocourse09/webinar/dropper"
	"github.com/sarff/prjctr-golang_beginning/gocourse09/webinar/drug"
	"github.com/sarff/prjctr-golang_beginning/gocourse09/webinar/patient"
)

func main() {
	patient1 := &patient.Patient{}
	drugControl := &drug.Drug{}
	dropperForPatient := &dropper.Dropper{
		CheckerPatient: patient1,
		ControllerDrug: drugControl,
	}
	dropperForPatient.React()
}
