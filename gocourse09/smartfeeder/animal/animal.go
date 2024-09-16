package animal

type Type string

const (
	Bear  Type = "Bear"
	Cow   Type = "Cow"
	Tiger Type = "Tiger"
)

type Animal struct {
	Type  Type
	Count int
}
