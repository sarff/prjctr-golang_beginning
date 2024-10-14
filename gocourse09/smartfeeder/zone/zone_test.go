package zone

import (
	"reflect"
	"testing"

	"github.com/sarff/prjctr-golang_beginning/gocourse09/smartfeeder/animal"
)

func TestZone_CheckZone(t *testing.T) {
	wantZone := &Zone{
		AnimalsInZone: []animal.Animal{
			{Type: animal.Tiger, Count: 3},
			{Type: animal.Cow, Count: 2},
			{Type: animal.Bear, Count: 2},
			{Type: animal.Panda, Count: 22},
			{Type: animal.Tiger, Count: 3},
			{Type: animal.Cow, Count: 2},
			{Type: animal.Bear, Count: 2},
			{Type: animal.Panda, Count: 22},
		},
	}

	newZone := &Zone{
		AnimalsInZone: []animal.Animal{
			{Type: animal.Tiger, Count: 3},
			{Type: animal.Cow, Count: 2},
			{Type: animal.Bear, Count: 2},
			{Type: animal.Panda, Count: 22},
		},
	}

	getZone := newZone.AddAnimals()
	if !reflect.DeepEqual(wantZone, getZone) {
		t.Errorf("AddAnimals returned %+v, want %+v", getZone, wantZone)
	}
}
