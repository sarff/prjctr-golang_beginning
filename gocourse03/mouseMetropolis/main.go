package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

/*
У зоопарку є окрема кімната «малі гризуни», де знаходиться кілька сімейств кількох видів гризунів. У кожного є чіп з його ID і видом.
Для них у цій кімнаті побудована мережа лабіринтів. У цьому лабіринті є один датчик, через який впродовж дня проходять сотні гризунів,
а до ночі вони розходяться по своїх куточках.
Для працівників зоопарку треба зробити програму, яка буде давати звіт, де знаходився кожний гризун на початку дня і де зупинився
наприкінці, а також зберігати історію його рухів через датчик.
У програмі може бути тип/структура для певного виду гризуна, набори (слайси) цих гризунів, і слайс для зберігання руху гризунів. Також стан сектору лабіринту: скільки і які гризуни знаходиться там у певний момент.
Треба симулювати рух гризунів через головний датчик і датчики секторів. Виконати операції видалення, пошуку і додавання в слайд.
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

type Mover interface {
	move()
	addMovementToHistory()
}

type Rodenter interface {
	searchMovementByRodentID()
	presenceOfRodentsInSector()
}

const (
	hoursHistory            = 1
	rodentCount             = 15
	rat          RodentType = "rat"
	mouse        RodentType = "mouse"
	sector1      Sector     = "sector1"
	sector2      Sector     = "sector2"
	sector3      Sector     = "sector3"
	sector4      Sector     = "sector4"
	sector5      Sector     = "sector5"
	sensor       Sector     = "sensor"
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

func newMovement() *Movement {
	return &Movement{
		Time:   time.Now(),
		FromTo: FromTo{},
	}
}

func (r *Rodent) move(movement Movement) {
	if r.History == nil {
		r.History = make([]Movement, hoursHistory)
	}

	r.History = append(r.History, movement)
}

func (r *Rodent) addMovementToHistory(movement Movement) {

}

func (r *Rodent) searchMovementByRodentID(id int) {

}

func (s Sector) presenceOfRodentsInSector(sector Sector, time time.Time) *Rodent {

	return &Rodent{}
}

func main() {
	//TODO кількість днів які будемо вести запис руху
	// кожну годину гризун рухається між секторами
	// дивимось в якому він секторі і рухаємо його в інший сектор рандомно - записуємо рух
	movement := newMovement()
	for d := 0; d < hoursHistory; d++ {
		for i := 0; i < rodentCount; i++ {
			rodent := newRodent(i)
			rodent.move(*movement)
			//fmt.Println(movement)
			fmt.Println(rodent)
		}
	}
}
