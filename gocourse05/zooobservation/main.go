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
)

func main() {
	controller := camera.Controller{
		DayCamera:   camera.DayCamera{},
		NightCamera: camera.NightCamera{},
	}
	tiger := animal.Animal{
		ID:      1,
		Species: "tiger",
	}
	bear := animal.Animal{
		ID:      2,
		Species: "bear",
	}
	var history []camera.HistoryItem
	var err error

	directions := [...]camera.Direction{camera.Left, camera.Right, camera.Top, camera.Bottom}
	for range 10 {

		history, err = controller.Move(tiger, directions[rand.IntN(len(directions))], history)
		if err != nil {
			fmt.Println(err)
		}
		history, err = controller.Move(bear, directions[rand.IntN(len(directions))], history)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Motion history to be transmitted to the server:")
	for _, d := range history {
		fmt.Printf("Time: %s || Direction: %s || ID: %d\n", d.Time.Format(time.RFC3339), d.Direction, d.ID)
	}
}
