package camera

import (
	"fmt"
	"github.com/sarff/prjctr-golang_beginning/gocourse05/zooobservation/common"
	"time"
)

type DayCamera struct {
	Screenshot string
}

type NightCamera struct {
	Screenshot string
}

type Camera interface {
	DetectMovement(direction common.Direction, historyItems []common.HistoryItem, animalID int) ([]common.HistoryItem, error)
	SaveToServer(historyItems []common.HistoryItem) error
}

func (d *DayCamera) DetectMovement(direction common.Direction, historyItems []common.HistoryItem, animalID int) ([]common.HistoryItem, error) {
	historyItems = moveToFront(direction, historyItems, animalID)
	return historyItems, d.SaveToServer(historyItems)
}

func (n *NightCamera) DetectMovement(direction common.Direction, historyItems []common.HistoryItem, animalID int) ([]common.HistoryItem, error) {
	historyItems = moveToFront(direction, historyItems, animalID)
	return historyItems, n.SaveToServer(historyItems)
}

func (d *DayCamera) SaveToServer(historyItems []common.HistoryItem) error {
	fmt.Println("Simulation: DayCamera history saved:", historyItems)
	return nil
}

func (n *NightCamera) SaveToServer(historyItems []common.HistoryItem) error {
	fmt.Println("Simulation: NightCamera history saved:", historyItems)
	return nil
}

func moveToFront(direction common.Direction, historyItems []common.HistoryItem, animalID int) []common.HistoryItem {
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
	historyItems = append(historyItems, common.HistoryItem{
		Time:      time.Now(),
		Direction: prev,
		AnimalID:  animalID,
	},
	)
	return historyItems
}
