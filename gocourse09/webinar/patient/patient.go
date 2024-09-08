package patient

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
	return ConditionGood
}
