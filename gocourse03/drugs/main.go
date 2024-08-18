/*
Треба відсортувати ящик ліків за датою виробництва.
Застарілі
Нові
яким півроку
*/
package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"time"
)

type Drugs struct {
	Name         string
	Manufactured time.Time
}

func main() {
	fake := gofakeit.New(0)
	dateRangeAt := time.Now().AddDate(-2, 0, 0)
	dateRangeTo := time.Now().AddDate(0, 0, 1)
	boxOfDrugs := []Drugs{
		{"Acetaminophen", fake.DateRange(dateRangeAt, dateRangeTo)},
		{"Cyclobenzaprine", fake.DateRange(dateRangeAt, dateRangeTo)},
		{"Lofexidine", fake.DateRange(dateRangeAt, dateRangeTo)},
		{"Pantoprazole", fake.DateRange(dateRangeAt, dateRangeTo)},
		{"Methadone", fake.DateRange(dateRangeAt, dateRangeTo)},
		{"Farxiga", fake.DateRange(dateRangeAt, dateRangeTo)},
		{"Brilinta", fake.DateRange(dateRangeAt, dateRangeTo)},
	}

	var oldDrugs = make([]Drugs, 0)
	var newDrugs = make([]Drugs, 0)
	var sixMonthDrugs = make([]Drugs, 0)
	sixMonthsAgo := time.Now().AddDate(0, -6, 0)

	for i, drug := range boxOfDrugs {
		switch {
		case drug.Manufactured.After(sixMonthsAgo):
			newDrugs = append(newDrugs, boxOfDrugs[i])
		case drug.Manufactured.Before(sixMonthsAgo) && drug.Manufactured.After(sixMonthsAgo.AddDate(0, -6, 0)):
			sixMonthDrugs = append(sixMonthDrugs, boxOfDrugs[i])
		default:
			oldDrugs = append(oldDrugs, boxOfDrugs[i])
		}
	}

	fmt.Println("List of medicines to be sorted: ", boxOfDrugs)
	fmt.Println("Today's date: ", time.Now())
	fmt.Println("list of medicines after sorting: ")
	fmt.Println("For sale: ", newDrugs)
	fmt.Println("For use: ", sixMonthDrugs)
	fmt.Println("For disposal: ", oldDrugs)
}
