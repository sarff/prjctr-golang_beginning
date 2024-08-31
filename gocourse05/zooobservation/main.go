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
	DetectMovement(direction Direction, history *History, animalID int) error
	SaveToServer(history *History) error
}

type DayCamera struct {
	screenshot string
}

type NightCamera struct {
	screenshot string
}

func (d *DayCamera) DetectMovement(direction Direction, history *History, animalID int) error {
	*history = append(*history, struct {
		time      time.Time
		direction Direction
		animalID  int
	}{
		time:      time.Now(),
		direction: direction,
		animalID:  animalID,
	})
	return d.SaveToServer(history)
}

func (n *NightCamera) DetectMovement(direction Direction, history *History, animalID int) error {
	*history = append(*history, struct {
		time      time.Time
		direction Direction
		animalID  int
	}{
		time:      time.Now(),
		direction: direction,
		animalID:  animalID,
	})
	return n.SaveToServer(history)
}

func (d *DayCamera) SaveToServer(history *History) error {
	fmt.Println("DayCamera: History saved:", *history)
	return nil
}

func (n *NightCamera) SaveToServer(history *History) error {
	fmt.Println("NightCamera: History saved:", *history)
	return nil
}

type Animal struct {
	id      int
	cam     Camera
	species string
}

func (t *Animal) Move(direction Direction, history *History) error {
	return t.cam.DetectMovement(direction, history, t.id)
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
		cam:     &dayCamera,
		species: "tiger",
	}
	bear := Animal{
		id:      2,
		cam:     &nightCamera,
		species: "bear",
	}
	history := &History{}

	direct := [...]Direction{left, right, top, bottom}
	for i := 0; i < 10; i++ {
		err := tiger.Move(direct[rand.IntN(len(direct))], history)
		if err != nil {
			fmt.Println(err)
		}
		err = bear.Move(direct[rand.IntN(len(direct))], history)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Motion history to be transmitted to the server:")
	for _, d := range *history {
		fmt.Printf("Time: %s || Direction: %s || AnimalID: %d\n", d.time.Format(time.RFC3339), d.direction, d.animalID)
	}
}
