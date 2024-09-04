package camera

import (
	"fmt"
	"time"
)

type Direction string

const (
	Left   Direction = "Left"
	Right  Direction = "Right"
	Top    Direction = "Top"
	Bottom Direction = "Bottom"
)

type DayCamera struct {
	Screenshot string
}

type NightCamera struct {
	Screenshot string
}

type HistoryItem struct {
	Time      time.Time
	Direction Direction
	AnimalID  int
}

type Camera interface {
	DetectMovement(direction Direction, historyItems []HistoryItem, animalID int) ([]HistoryItem, error)
	SaveToServer(historyItems []HistoryItem) error
}

func (d *DayCamera) DetectMovement(direction Direction, historyItems []HistoryItem, animalID int) ([]HistoryItem, error) {
	historyItems = moveToFront(direction, historyItems, animalID)
	return historyItems, d.SaveToServer(historyItems)
}

func (n *NightCamera) DetectMovement(direction Direction, historyItems []HistoryItem, animalID int) ([]HistoryItem, error) {
	historyItems = moveToFront(direction, historyItems, animalID)
	return historyItems, n.SaveToServer(historyItems)
}

func (d *DayCamera) SaveToServer(historyItems []HistoryItem) error {
	fmt.Println("Simulation: DayCamera history saved:", historyItems)
	return nil
}

func (n *NightCamera) SaveToServer(historyItems []HistoryItem) error {
	fmt.Println("Simulation: NightCamera history saved:", historyItems)
	return nil
}

func moveToFront(direction Direction, historyItems []HistoryItem, animalID int) []HistoryItem {
	prev := direction
	for i, elem := range historyItems {
		switch {
		case i == 0:
			historyItems[0].Direction = direction
			prev = elem.Direction
		case elem.Direction == direction:
			historyItems[i].Direction = prev
			return historyItems
		default:
			historyItems[i].Direction = prev
			prev = elem.Direction
		}
	}
	historyItems = append(historyItems, HistoryItem{
		Time:      time.Now(),
		Direction: prev,
		AnimalID:  animalID,
	},
	)
	return historyItems
}
