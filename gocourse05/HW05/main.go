/*
Треба зробити бекенд для сервера, на який надходять нічні зображення з різних камер спостереження з датчиком руху,
розвішаних по всьому зоопарку. У зоопарку існує кулька типів камер. Деякі камери працюють із зовнішнім світлом, інші —
в нічному режимі. Треба обробляти дані з різних джерел (типів камер), зберігати в памʼяті історію подій і передавати
єдиний уніфікований запит на інший сервер. Відповідно, треба зробити кілька типів (структур) які відповідають своїм
реальним камерам, і декілька інтерфейсів, із якими працює програма. Використовувати контракти і обробляти можливі помилки.
Також треба написати тести для позитивних і негативних випадків роботи функцій, які оброблюють дані з камер.
Тут «сервер» — умовна назва для нашої програми.
Даними може виступати рух певної тварини. Наприклад: тигр, пішов ліворуч; ведмідь, стоїть.
*/
//TODO: Написати коменти і тести
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
	TakeScreenshot(direction Direction, history *History, animalID int) error
	SaveToServer(direction Direction, history *History, animalID int) error
}

type DayCamera struct {
	screenshot string
}

type NightCamera struct {
	screenshot string
}

func (d *DayCamera) TakeScreenshot(direction Direction, history *History, animalID int) error {
	return d.SaveToServer(direction, history, animalID)
}

func (n *NightCamera) TakeScreenshot(direction Direction, history *History, animalID int) error {
	return n.SaveToServer(direction, history, animalID)

}

func (d *DayCamera) SaveToServer(direction Direction, history *History, animalID int) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Have a panic while saving to the server %v. Screen: %s\n", err, d.screenshot)
		}
	}()
	*history = append(*history, struct {
		time      time.Time
		direction Direction
		animalID  int
	}{time: time.Now().UTC(), direction: direction, animalID: animalID})

	return nil
}

func (n *NightCamera) SaveToServer(direction Direction, history *History, animalID int) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Have a panic while saving to the server %v. Screen: %s\n", err, n.screenshot)
		}
	}()

	*history = append(*history, struct {
		time      time.Time
		direction Direction
		animalID  int
	}{
		time:      time.Now(),
		direction: direction,
		animalID:  animalID,
	})

	return nil
}

type Movement interface {
	Move(direction Direction)
}

type Bear struct {
	ID  int
	cam NightCamera
}

type Tiger struct {
	ID  int
	cam DayCamera
}

func (t *Tiger) Move(direction Direction, history *History) error {
	return t.cam.TakeScreenshot(direction, history, t.ID)
}

func (b *Bear) Move(direction Direction, history *History) error {
	return b.cam.TakeScreenshot(direction, history, b.ID)
}

func main() {
	dayCam := DayCamera{
		screenshot: "day_screenshot.png",
	}
	nightCam := NightCamera{
		screenshot: "night_screenshot.png",
	}
	tiger := Tiger{
		ID:  1,
		cam: dayCam,
	}
	bear := Bear{
		ID:  2,
		cam: nightCam,
	}
	history := &History{}
	err := error(nil)

	direct := [...]Direction{left, right, top, bottom}
	for i := 0; i < 10; i++ {
		err = tiger.Move(direct[rand.IntN(len(direct))], history)
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
