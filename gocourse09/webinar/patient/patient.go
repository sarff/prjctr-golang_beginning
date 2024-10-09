package patient

import "math/rand/v2"

type ConditionPatient string

const (
	ConditionCritical ConditionPatient = "Critical"
	ConditionNormal   ConditionPatient = "Normal"
	ConditionGood     ConditionPatient = "Good"
)

type Patient struct {
	CheckerPatient CheckerPatient
}

type CheckerPatient interface {
	CheckCondition() ConditionPatient
}

func (p *Patient) CheckCondition() ConditionPatient {
	conditions := [...]ConditionPatient{ConditionCritical, ConditionNormal, ConditionGood}
	randConditions := conditions[rand.IntN(len(conditions))]
	return randConditions
}
