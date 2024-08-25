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
		Zoo      = &Zoo{Areas: lookupAreas}
		wantName = "Horse1"
	)

	_, _, animal := Zoo.Lookup(wantName)
	if animal != nil && animal.Name != wantName {
		t.Errorf("expected an animal named %s, but got %s", wantName, animal.Name)
	}
}

func TestClean(t *testing.T) {
	var (
		testZoo  = &Zoo{Areas: lookupAreas}
		Cleaning = &Zoo{Areas: testsCleaning}
		Name     = "Horse1"
	)

	AnimalArea, AnimalSector, Animal := testZoo.Lookup(Name)
	Animal.Clean(AnimalArea.Name, AnimalSector.ID, *testZoo)
	if !reflect.DeepEqual(testZoo, Cleaning) {
		t.Errorf("Output \n%v not equal to expected \n%v", testZoo, Cleaning)
	}
}

func TestMigration(t *testing.T) {
	var (
		testZoo      = &Zoo{Areas: lookupAreas}
		testMigrated = &Zoo{Areas: testsMigrated}
		Name         = "Horse1"
		ToArea       = "tech room"
		ToSector     = 5
	)

	AnimalArea, AnimalSector, Animal := testZoo.Lookup(Name)
	_ = testZoo.Migration(AnimalArea.Name, ToArea, AnimalSector.ID, ToSector, *Animal)
	if !reflect.DeepEqual(testZoo, testMigrated) {
		t.Errorf("Output \n%v not equal to expected \n%v", testMigrated, testZoo)
	}
}

func BenchmarkDeleteAnimal(b *testing.B) {
	var (
		Zoo         = &Zoo{Areas: lookupAreas}
		idForDelete int
	)

	for n := 0; n < b.N; n++ {
		Zoo.deleteAnimal(idForDelete)
	}
}
