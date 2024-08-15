package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"time"
)

/*
Треба відсортувати ящик ліків за датою виробництва.
Застарілі
Нові
яким півроку
*/

type Drugs struct {
	Name             string
	ManufacturerDate time.Time
}

func main() {
	fake := gofakeit.New(0)

	var oldDrugs = make([]Drugs, 0)
	var newDrugs = make([]Drugs, 0)
	var sixMonthDrugs = make([]Drugs, 0)

	boxOfDrugs := []Drugs{
		{"Acetaminophen", fake.DateRange(time.Now().AddDate(-1, 0, 0), time.Now().AddDate(0, 0, 1))},
		{"Cyclobenzaprine", fake.DateRange(time.Now().AddDate(-1, 0, 0), time.Now().AddDate(0, 0, 1))},
		{"Lofexidine", fake.DateRange(time.Now().AddDate(-1, 0, 0), time.Now().AddDate(0, 0, 1))},
		{"Pantoprazole", fake.DateRange(time.Now().AddDate(-1, 0, 0), time.Now().AddDate(0, 0, 1))},
		{"Methadone", fake.DateRange(time.Now().AddDate(-1, 0, 0), time.Now().AddDate(0, 0, 1))},
		{"Farxiga", fake.DateRange(time.Now().AddDate(-1, 0, 0), time.Now().AddDate(0, 0, 1))},
		{"Brilinta", fake.DateRange(time.Now().AddDate(-1, 0, 0), time.Now().AddDate(0, 0, 1))},
	}

	for i, d := range boxOfDrugs {
		dayDiff := int(time.Now().Sub(d.ManufacturerDate) / (time.Hour * 24))
		switch {
		case dayDiff < 70:
			newDrugs = append(newDrugs, boxOfDrugs[i])
		case dayDiff > 70 && dayDiff < 200:
			sixMonthDrugs = append(sixMonthDrugs, boxOfDrugs[i])
		default:
			oldDrugs = append(oldDrugs, boxOfDrugs[i])
		}
	}

	fmt.Println(newDrugs)
	fmt.Println(sixMonthDrugs)
	fmt.Println(oldDrugs)

}
