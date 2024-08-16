package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"math/rand/v2"
	"time"
)

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
type (
	Sector     string
	RodentType string
	FromTo     [2]Sector

	Movement struct {
		Time time.Time
		FromTo
	}

	Rodent struct {
		ID      int
		Type    RodentType
		History []Movement
	}
)

const (
	rodentCount            = 15
	rat         RodentType = "rat"
	mouse       RodentType = "mouse"
	sector1     Sector     = "sector1"
	sector2     Sector     = "sector2"
	sector3     Sector     = "sector3"
	sector4     Sector     = "sector4"
	sector5     Sector     = "sector5"
	sensor      Sector     = "sensor"
)

func newRodent(id int) *Rodent {
	var typeRodent RodentType

	if rand.N(2) == 1 {
		typeRodent = rat
	} else {
		typeRodent = mouse
	}
	n := &Rodent{ID: id + 1, Type: typeRodent, History: nil}
	return n
}

func (r *Rodent) addHistory(movement Movement) {
	r.History = append(r.History, movement)
}

func (r *Rodent) delHistory(t time.Time) {
	for i, m := range r.History {
		if m.Time == t {
			fmt.Println("time found and successfully deleted: ", m)
			r.History = append(r.History[:i], r.History[i+1:]...)
		}
	}
}

func (r *Rodent) findHistory() {
	fmt.Println("move history: ", r.History)
}

func (r *Rodent) moveRodent(count int) {
	var from, to Sector
	if count == 0 {
		count++
	}

	for i := 0; i < count; i++ {
		if r.History == nil {
			from = sector1
			to = returnRandSector(from)
		} else {
			from = r.History[len(r.History)-1].FromTo[1]
			to = returnRandSector(from)
		}

		roundTime := time.Now().Add(time.Hour * time.Duration(i)).Truncate(time.Hour)
		movement := Movement{

			Time: roundTime,
			FromTo: FromTo{
				from,
				to,
			},
		}

		r.addHistory(movement)
	}
}

func returnRandSector(from Sector) Sector {
	fake := gofakeit.New(0)
	randSector := fake.RandomString([]string{string(sector1), string(sector2), string(sector3), string(sector4), string(sensor)})
	if Sector(randSector) != from {
		return Sector(randSector)
	} else if Sector(randSector) != sensor {
		return sensor
	}
	return sector5
}

func (r *Rodent) startSector() Sector {
	return r.History[0].FromTo[0]
}

func (r *Rodent) finishSector() Sector {
	return r.History[len(r.History)-1].FromTo[1]
}

func returnRodentInSector(time time.Time, rodents []Rodent, targetSector Sector) []int {
	var report = make([]int, 0)
	for _, rodent := range rodents {
		for _, history := range rodent.History {
			if history.Time == time && targetSector == history.FromTo[0] {
				report = append(report, rodent.ID)
			}
		}
	}
	return report
}

func main() {
	var rodents []Rodent

	for i := 0; i < rodentCount; i++ {
		rodent := newRodent(i)
		rodent.moveRodent(rand.IntN(10))
		rodents = append(rodents, *rodent)
	}
	fmt.Println(rodents)

	// delete time History by time, for ID = 5
	rodents[5].delHistory(time.Now().Add(time.Hour * time.Duration(1)).Truncate(time.Hour))

	// find history by ID = 4
	rodents[4].findHistory()

	// where the rodent was at the beginning of the day and where it stopped at the end of the day
	fmt.Println("Start Sector:", rodents[4].startSector(), "|| Finish Sector: ", rodents[4].finishSector())

	// report on the sector at the specified time
	filterTime := time.Now().Add(time.Hour * time.Duration(2)).Truncate(time.Hour)
	targetSector := sector1
	report := returnRodentInSector(filterTime, rodents, targetSector)
	if len(report) > 0 {
		fmt.Printf("the following rodentsID were in the %s at the time %s: %v", targetSector, filterTime, report)
	} else {
		fmt.Println("there were no rodents in this sector at this time")
	}
}
