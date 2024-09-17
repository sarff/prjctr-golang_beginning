package animal

type Type string

const (
	Bear  Type = "Bear"
	Cow   Type = "Cow"
	Tiger Type = "Tiger"
	Panda Type = "Panda"
)

type Animal struct {
	Type  Type
	Count int
}
