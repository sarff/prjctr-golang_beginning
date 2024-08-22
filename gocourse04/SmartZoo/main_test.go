package main

import (
	"reflect"
	"testing"
)

var tests = map[string]Area{
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
		testZoo  = &Zoo{Areas: tests}
		testName = "Horse1"
	)

	_, _, animal := testZoo.Lookup(testName)
	if animal != nil && animal.Name != testName {
		t.Errorf(`function returned a different object. We expected an object with the name: %s, 
						 but we got an object with the name %s`, testName, animal.Name)
	}
}

func TestCleaning(t *testing.T) {
	var (
		testZoo      = &Zoo{Areas: tests}
		testCleaning = &Zoo{Areas: testsCleaning}
		testName     = "Horse1"
	)

	testAnimalArea, testAnimalSector, testAnimal := testZoo.Lookup(testName)
	testAnimal.Clean(testAnimalArea.Name, testAnimalSector.ID, *testZoo)
	if !reflect.DeepEqual(testZoo, testCleaning) {
		t.Errorf("Output \n%v not equal to expected \n%v", testZoo, testCleaning)
	}
}

func TestMigration(t *testing.T) {
	var (
		testZoo      = &Zoo{Areas: tests}
		testMigrated = &Zoo{Areas: testsMigrated}
		testName     = "Horse1"
		testToArea   = "tech room"
		testToSector = 5
	)

	testAnimalArea, testAnimalSector, testAnimal := testZoo.Lookup(testName)
	_ = testZoo.Migration(testAnimalArea.Name, testToArea, testAnimalSector.ID, testToSector, *testAnimal)
	if !reflect.DeepEqual(testZoo, testMigrated) {
		t.Errorf("Output \n%v not equal to expected \n%v", testMigrated, testZoo)
	}
}

func BenchmarkDelAnimal(b *testing.B) {
	var (
		testZoo     = &Zoo{Areas: tests}
		idForDelete int
	)

	for n := 0; n < b.N; n++ {
		testZoo.deleteAnimal(idForDelete)
	}
}
