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

// Для PR descr
// Створюємо обʼєкт patient1 (не можемо назвати patient так як є імпортований модуль)
// Обʼєкт контроль ліків drugControl
// Обʼєкт крапильниці dropperForPatient (приймає patient1 і drugControl)
// Далі функція React() - яка змішує потрібні ліки
func main() {
	patient1 := &patient.Patient{}
	drugControl := &drug.Drug{}
	dropperForPatient := &dropper.Dropper{
		CheckerPatient: patient1,
		ControllerDrug: drugControl,
	}
	// TODO: додат виведення в консоль опис... який пацієнт, який стан
	dropperForPatient.React()
}
