/*
Треба зробити бекенд для сервера, на який надходять нічні зображення з різних камер спостереження з датчиком руху,
розвішаних по всьому зоопарку. У зоопарку існує кілька типів камер. Деякі камери працюють із зовнішнім світлом, інші —
в нічному режимі. Треба обробляти дані з різних джерел (типів камер), зберігати в памʼяті історію подій і передавати
єдиний уніфікований запит на інший сервер. Відповідно, треба зробити кілька типів (структур) які відповідають своїм
реальним камерам, і декілька інтерфейсів, із якими працює програма. Використовувати контракти й обробляти можливі помилки.
Також треба написати тести для позитивних і негативних випадків роботи функцій, які оброблюють дані з камер.
Тут «сервер» — умовна назва для нашої програми.
Даними може виступати рух певної тварини. Наприклад: тигр, пішов ліворуч; ведмідь, стоїть.
*/
package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type Direction string

const (
	left   Direction = "left"
	right  Direction = "right"
	top    Direction = "top"
	bottom Direction = "bottom"
)

type HistoryItem struct {
	time      time.Time
	direction Direction
	animalID  int
}

type History []HistoryItem

type Camera interface {
	DetectMovement(direction Direction, historyItems []HistoryItem, animalID int) ([]HistoryItem, error)
	SaveToServer(historyItems []HistoryItem) error
}

type DayCamera struct {
	screenshot string
}

type NightCamera struct {
	screenshot string
}

func moveToFront(direction Direction, historyItems []HistoryItem, animalID int) []HistoryItem {
	prev := direction
	for i, elem := range historyItems {
		switch {
		case i == 0:
			historyItems[0].direction = direction
			prev = elem.direction
		case elem.direction == direction:
			historyItems[i].direction = prev
			return historyItems
		default:
			historyItems[i].direction = prev
			prev = elem.direction
		}
	}
	historyItems = append(historyItems, HistoryItem{
		time:      time.Now(),
		direction: prev,
		animalID:  animalID,
	},
	)
	return historyItems
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
	fmt.Println("DayCamera: History saved:", historyItems)
	return nil
}

func (n *NightCamera) SaveToServer(historyItems []HistoryItem) error {
	fmt.Println("NightCamera: History saved:", historyItems)
	return nil
}

type Animal struct {
	id      int
	camera  Camera
	species string
}

func (t *Animal) Move(direction Direction, historyItems []HistoryItem) ([]HistoryItem, error) {
	return t.camera.DetectMovement(direction, historyItems, t.id)
}

func main() {
	dayCamera := DayCamera{
		screenshot: "day_screenshot.png",
	}
	nightCamera := NightCamera{
		screenshot: "night_screenshot.png",
	}
	tiger := Animal{
		id:      1,
		camera:  &dayCamera,
		species: "tiger",
	}
	bear := Animal{
		id:      2,
		camera:  &nightCamera,
		species: "bear",
	}
	var history []HistoryItem
	var err error

	directions := [...]Direction{left, right, top, bottom}
	for range 10 {
		history, err = tiger.Move(directions[rand.IntN(len(directions))], history)
		if err != nil {
			fmt.Println(err)
		}
		history, err = bear.Move(directions[rand.IntN(len(directions))], history)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Motion history to be transmitted to the server:")
	for _, d := range history {
		fmt.Printf("Time: %s || Direction: %s || AnimalID: %d\n", d.time.Format(time.RFC3339), d.direction, d.animalID)
	}
}
