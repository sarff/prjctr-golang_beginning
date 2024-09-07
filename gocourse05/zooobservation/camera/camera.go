package camera

import (
	"fmt"
	"time"

	"github.com/sarff/prjctr-golang_beginning/gocourse05/zooobservation/animal"
)

type Direction string

const (
	Left   Direction = "Left"
	Right  Direction = "Right"
	Top    Direction = "Top"
	Bottom Direction = "Bottom"
)

type HistoryItem struct {
	Time      time.Time
	Direction Direction
	ID        int
}

type Camera interface {
	DetectMovement(direction Direction, historyItems []HistoryItem, animalID int) ([]HistoryItem, error)
	SaveToServer(historyItems []HistoryItem) error
}

type DayCamera struct{}

func (d *DayCamera) DetectMovement(direction Direction, historyItems []HistoryItem, animalID int) ([]HistoryItem, error) {
	historyItems = moveToFront(direction, historyItems, animalID)
	return historyItems, d.SaveToServer(historyItems)
}

func (d *DayCamera) SaveToServer(historyItems []HistoryItem) error {
	fmt.Println("Simulation: DayCamera history saved:", historyItems)
	return nil
}

type NightCamera struct{}

func (n *NightCamera) DetectMovement(direction Direction, historyItems []HistoryItem, animalID int) ([]HistoryItem, error) {
	historyItems = moveToFront(direction, historyItems, animalID)
	return historyItems, n.SaveToServer(historyItems)
}

func (n *NightCamera) SaveToServer(historyItems []HistoryItem) error {
	fmt.Println("Simulation: NightCamera history saved:", historyItems)
	return nil
}

type Controller struct {
	DayCamera   DayCamera
	NightCamera NightCamera
}

func (c *Controller) Move(animal animal.Animal, direction Direction, historyItems []HistoryItem) ([]HistoryItem, error) {
	var camera Camera
	camera = &c.NightCamera
	if animal.Species == "tiger" {
		camera = &c.DayCamera
	}
	return camera.DetectMovement(direction, historyItems, animal.ID)
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
		ID:        animalID,
	},
	)
	return historyItems
}
