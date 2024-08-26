/*
Беремо той самий ящик ліків із минулого заняття, але тепер дані в мапах.

На вході маємо великий ящик із ліками, на виході — три маленьких і пустий початковий.
На вимогу логіста, засовуємо три нові ящики в один великий, який вміє шукати в цих трьох маленьких.
*/
package main

import (
	"fmt"
	"time"
)

type Medicine struct {
	Name         string
	Manufactured time.Time
}

type Box struct {
	ToDiscard map[string]Medicine
	ForSale   map[string]Medicine
	ToUse     map[string]Medicine // map, where key is medicine name, value is Medicine
}

func main() {
	today := time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC)

	medicines := map[string]Medicine{
		"MedicineA": {Name: "MedicineA", Manufactured: time.Date(2023, time.July, 1, 0, 0, 0, 0, time.UTC)},
		"MedicineB": {Name: "MedicineB", Manufactured: time.Date(2022, time.August, 1, 0, 0, 0, 0, time.UTC)},
		"MedicineC": {Name: "MedicineC", Manufactured: time.Date(2023, time.October, 1, 0, 0, 0, 0, time.UTC)},
	}

	box := NewBox()
	box.FilterMedicines(medicines, today)

	if med, category := box.FindMedicine("MedicineA"); med != nil {
		fmt.Printf("Ліки '%s' знайдені у категорії '%s'\n", med.Name, category)
	} else {
		fmt.Println("Ліки 'MedicineA' не знайдені.")
	}
}

func NewBox() *Box {
	return &Box{
		ToDiscard: make(map[string]Medicine),
		ForSale:   make(map[string]Medicine),
		ToUse:     make(map[string]Medicine),
	}
}

func (b *Box) FilterMedicines(medicines map[string]Medicine, today time.Time) {
	sixMonthsAgo := today.AddDate(0, -6, 0)
	for name, medicine := range medicines {
		if medicine.Manufactured.After(sixMonthsAgo) {
			b.ForSale[name] = medicine
		} else if medicine.Manufactured.Before(sixMonthsAgo) && medicine.Manufactured.After(sixMonthsAgo.AddDate(0, -6, 0)) {
			b.ToUse[name] = medicine
		} else {
			b.ToDiscard[name] = medicine
		}
	}
}

func (b *Box) FindMedicine(s string) (*Medicine, string) {
	if val, ok := b.ToDiscard[s]; ok {
		return &val, "ToDiscard"
	}
	if val, ok := b.ForSale[s]; ok {
		return &val, "ForSale"
	}
	if val, ok := b.ToUse[s]; ok {
		return &val, "ToUse"
	}
	return nil, ""
}
