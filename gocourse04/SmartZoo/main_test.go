package main

import (
	"reflect"
	"testing"
)

var lookupAreas = map[string]Area{
	"Hoofed": {
		Name: "hoofed", Sectors: map[string]Sector{
			"Horses": {
				ID: 1,
				Animals: []Animal{
					{ID: 1, Name: "Horse1"},
					{ID: 2, Name: "Horse2"},
				},
			},
		},
	},
	"Birds": {
		Name: "birds", Sectors: map[string]Sector{
			"Doves": {
				ID: 3,
				Animals: []Animal{
					{ID: 3, Name: "Dove1"},
					{ID: 4, Name: "Dove2"},
				},
			},
		},
	},
	"Techroom": {
		Name: "tech room", Sectors: map[string]Sector{
			"Dining room": {
				ID: 5,
				Animals: []Animal{
					{ID: 5, Name: "Hawk4"},
				},
			},
			"Bath room": {
				ID: 6,
				Animals: []Animal{
					{ID: 6, Name: "Cow4"},
				},
			},
		},
	},
}

var testsCleaning = map[string]Area{
	"Hoofed": {
		Name: "hoofed", Sectors: map[string]Sector{
			"Horses": {
				ID: 1,
				Animals: []Animal{
					{ID: 2, Name: "Horse2"},
					{ID: 1, Name: "Horse1"},
				},
			},
		},
	},
	"Birds": {
		Name: "birds", Sectors: map[string]Sector{
			"Doves": {
				ID: 3,
				Animals: []Animal{
					{ID: 3, Name: "Dove1"},
					{ID: 4, Name: "Dove2"},
				},
			},
		},
	},
	"Techroom": {
		Name: "tech room", Sectors: map[string]Sector{
			"Dining room": {
				ID: 5,
				Animals: []Animal{
					{ID: 5, Name: "Hawk4"},
				},
			},
			"Bath room": {
				ID: 6,
				Animals: []Animal{
					{ID: 6, Name: "Cow4"},
				},
			},
		},
	},
}

var testsMigrated = map[string]Area{
	"Hoofed": {
		Name: "hoofed", Sectors: map[string]Sector{
			"Horses": {
				ID: 1,
				Animals: []Animal{
					{ID: 2, Name: "Horse2"},
				},
			},
		},
	},
	"Birds": {
		Name: "birds", Sectors: map[string]Sector{
			"Doves": {
				ID: 3,
				Animals: []Animal{
					{ID: 3, Name: "Dove1"},
					{ID: 4, Name: "Dove2"},
				},
			},
		},
	},
	"Techroom": {
		Name: "tech room", Sectors: map[string]Sector{
			"Dining room": {
				ID: 5,
				Animals: []Animal{
					{ID: 5, Name: "Hawk4"},
					{ID: 1, Name: "Horse1"},
				},
			},
			"Bath room": {
				ID: 6,
				Animals: []Animal{
					{ID: 6, Name: "Cow4"},
				},
			},
		},
	},
}

func TestLookup(t *testing.T) {
	var (
		zoo      = &Zoo{Areas: lookupAreas}
		wantName = "Horse1"
	)

	_, _, animal := zoo.Lookup(wantName)
	if animal != nil && animal.Name != wantName {
		t.Errorf("expected an animal named %s, but got %s", wantName, animal.Name)
	}
}

func TestClean(t *testing.T) {
	var (
		zoo      = &Zoo{Areas: lookupAreas}
		cleaning = &Zoo{Areas: testsCleaning}
		name     = "Horse1"
	)

	AnimalArea, AnimalSector, Animal := zoo.Lookup(name)
	Animal.Clean(AnimalArea.Name, AnimalSector.ID, *zoo)
	if !reflect.DeepEqual(zoo, cleaning) {
		t.Errorf("Output \n%v not equal to expected \n%v", zoo, cleaning)
	}
}

func TestMigration(t *testing.T) {
	var (
		zoo      = &Zoo{Areas: lookupAreas}
		migrated = &Zoo{Areas: testsMigrated}
		name     = "Horse1"
		toArea   = "tech room"
		toSector = 5
	)

	AnimalArea, AnimalSector, Animal := zoo.Lookup(name)
	_ = zoo.MoveAnimal(AnimalArea.Name, toArea, AnimalSector.ID, toSector, *Animal)
	if !reflect.DeepEqual(zoo, migrated) {
		t.Errorf("Output \n%v not equal to expected \n%v", migrated, zoo)
	}
}

func BenchmarkDeleteAnimal(b *testing.B) {
	var (
		zoo         = &Zoo{Areas: lookupAreas}
		idForDelete int
	)

	for n := 0; n < b.N; n++ {
		zoo.deleteAnimal(idForDelete)
	}
}
