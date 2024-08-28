/*
Треба відсортувати ящик ліків за датою виробництва.
Застарілі
Нові
яким півроку
*/
package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type Drugs struct {
	Name         string
	Manufactured time.Time
}

func main() {
	timeNow := time.Now()
	fake := gofakeit.New(0)
	fakeStart := timeNow.AddDate(-2, 0, 0)
	fakeEnd := timeNow.AddDate(0, 0, 1)
	boxOfDrugs := []Drugs{
		{"Acetaminophen", fake.DateRange(fakeStart, fakeEnd)},
		{"Cyclobenzaprine", fake.DateRange(fakeStart, fakeEnd)},
		{"Lofexidine", fake.DateRange(fakeStart, fakeEnd)},
		{"Pantoprazole", fake.DateRange(fakeStart, fakeEnd)},
		{"Methadone", fake.DateRange(fakeStart, fakeEnd)},
		{"Farxiga", fake.DateRange(fakeStart, fakeEnd)},
		{"Brilinta", fake.DateRange(fakeStart, fakeEnd)},
	}

	var oldDrugs []Drugs
	var newDrugs []Drugs
	var sixMonthDrugs []Drugs
	sixMonthsAgo := timeNow.AddDate(0, -6, 0)

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
	fmt.Println("Today's date: ", timeNow)
	fmt.Println("list of medicines after sorting: ")
	fmt.Println("For sale: ", newDrugs)
	fmt.Println("For use: ", sixMonthDrugs)
	fmt.Println("For disposal: ", oldDrugs)
}
