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

	"github.com/sarff/prjctr-golang_beginning/gocourse05/zooobservation/animal"
	"github.com/sarff/prjctr-golang_beginning/gocourse05/zooobservation/camera"
	"github.com/sarff/prjctr-golang_beginning/gocourse05/zooobservation/common"
)

func main() {
	dayCamera := camera.DayCamera{
		Screenshot: "day_screenshot.png",
	}
	nightCamera := camera.NightCamera{
		Screenshot: "night_screenshot.png",
	}
	tiger := animal.Animal{
		ID:      1,
		Camera:  &dayCamera,
		Species: "tiger",
	}
	bear := animal.Animal{
		ID:      2,
		Camera:  &nightCamera,
		Species: "bear",
	}
	var history []common.HistoryItem
	var err error

	directions := [...]common.Direction{common.Left, common.Right, common.Top, common.Bottom}
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
		fmt.Printf("Time: %s || Direction: %s || ID: %d\n", d.Time.Format(time.RFC3339), d.Direction, d.ID)
	}
}
