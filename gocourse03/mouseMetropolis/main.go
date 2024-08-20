/*
У зоопарку є окрема кімната «малі гризуни», де знаходиться кілька сімейств кількох видів гризунів. У кожного є чіп з
його ID і видом. Для них у цій кімнаті побудована мережа лабіринтів. У цьому лабіринті є один датчик, через який
впродовж дня проходять сотні гризунів, а до ночі вони розходяться по своїх куточках. Для працівників зоопарку треба
зробити програму, яка буде давати звіт, де знаходився кожний гризун на початку дня і де зупинився
наприкінці, а також зберігати історію його рухів через датчик. У програмі може бути тип/структура для певного виду
гризуна, набори (слайси) цих гризунів, і слайс для зберігання руху гризунів. Також стан сектору лабіринту: скільки
і які гризуни знаходиться там у певний момент. Треба симулювати рух гризунів через головний датчик і датчики секторів.
Виконати операції видалення, пошуку і додавання в слайд.
*/
package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type (
	Sector     string
	RodentType string
	FromTo     [2]Sector

	Movement struct {
		Time time.Time
		FromTo
	}

	MiceRoom struct {
		Rodents []Rodent
	}

	Rodent struct {
		ID      int
		Type    RodentType
		History []Movement
	}
)

const (
	rodentCount = 15

	rat   RodentType = "rat"
	mouse RodentType = "mouse"

	sectorA Sector = "sectorA"
	sectorB Sector = "sectorB"
	sectorC Sector = "sectorC"
	sectorD Sector = "sectorD"
	sectorE Sector = "sectorE"
)

// this constructor randomly fills the typeRodent
func newRodent(id int) *Rodent {
	typeRodent := mouse

	if rand.N(2) == 1 {
		typeRodent = rat
	}
	// make +1 because the ID cannot be 0
	return &Rodent{ID: id + 1, Type: typeRodent, History: nil}
}

func (r *Rodent) addMovement(movement Movement) {
	r.History = append(r.History, movement)
}

func (m *MiceRoom) delMovement(t time.Time) bool {
	delStatus := false
	for _, rodent := range m.Rodents {
		for i, m := range rodent.History {
			if m.Time == t {
				rodent.History = append(rodent.History[:i], rodent.History[i+1:]...)
				delStatus = true
			}
		}
	}
	return delStatus
}

func (m *MiceRoom) findHistory(id int) map[time.Time][]FromTo {
	history := make(map[time.Time][]FromTo)
	for _, rodent := range m.Rodents {
		if rodent.ID == id {
			for _, h := range rodent.History {
				history[h.Time] = append(history[h.Time], h.FromTo)
			}
		}
	}
	return history
}

func (r *Rodent) moveRodent(count int) {
	var from, to Sector
	if count == 0 {
		count++
	}

	for i := 0; i < count; i++ {
		if r.History == nil {
			from = sectorA
		} else {
			from = r.History[len(r.History)-1].FromTo[1]
		}
		to = randSectorExcluding(from)

		roundTime := time.Now().Add(time.Hour * time.Duration(i)).Truncate(time.Hour)
		movement := Movement{
			Time: roundTime,
			FromTo: FromTo{
				from,
				to,
			},
		}
		r.addMovement(movement)
	}
}

func randSectorExcluding(exclude Sector) Sector {
	fake := gofakeit.New(0)
	randSector := fake.RandomString([]string{string(sectorA), string(sectorB), string(sectorC), string(sectorD)})
	if Sector(randSector) != exclude {
		return Sector(randSector)
	}
	return sectorE
}

func (m *MiceRoom) startSector(ID int) Sector {
	for _, rodent := range m.Rodents {
		if rodent.ID == ID {
			return rodent.History[0].FromTo[0]
		}
	}
	return ""
}

func (m *MiceRoom) finishSector(ID int) Sector {
	for _, rodent := range m.Rodents {
		if rodent.ID == ID {
			return rodent.History[len(rodent.History)-1].FromTo[1]
		}
	}
	return ""
}

func (m *MiceRoom) returnRodentInSector(time time.Time, targetSector Sector) []int {
	report := make([]int, 0)
	for _, rodent := range m.Rodents {
		for _, history := range rodent.History {
			if history.Time == time && targetSector == history.FromTo[0] {
				report = append(report, rodent.ID)
			}
		}
	}
	return report
}

func main() {
	var miceRoom MiceRoom

	for i := 0; i < rodentCount; i++ {
		rodent := newRodent(i)
		rodent.moveRodent(rand.IntN(10))
		miceRoom.Rodents = append(miceRoom.Rodents, *rodent)
	}

	// find history by ID = 4
	rodentID := 4
	history := miceRoom.findHistory(rodentID)
	fmt.Printf("Move history for ID %d :\n", rodentID)
	for k, v := range history {
		fmt.Printf("-->   %s : %v \n", k, v)
	}

	// where the rodent was at the beginning of the day and where it stopped at the end of the day
	fmt.Println("Start and Finish sectors for RodentID:", rodentID)
	fmt.Println("Start Sector:", miceRoom.startSector(rodentID), "\nFinish Sector: ", miceRoom.finishSector(rodentID))

	// report on the sector at the specified time
	filterTime := time.Now().Add(time.Hour * time.Duration(2)).Truncate(time.Hour)
	targetSector := sectorA
	report := miceRoom.returnRodentInSector(filterTime, targetSector)
	if len(report) > 0 {
		fmt.Printf("The following rodentsID were in the %s at the time %s: %v\n", targetSector, filterTime, report)
	} else {
		fmt.Println("There were no rodents in this sector at this time\n")
	}

	// delete time History by time, for all rodents
	timeForDel := time.Now().Add(time.Hour * time.Duration(1)).Truncate(time.Hour)
	if miceRoom.delMovement(timeForDel) {
		fmt.Printf("All entries with `%s` time are deleted\n", timeForDel)
	}

}
