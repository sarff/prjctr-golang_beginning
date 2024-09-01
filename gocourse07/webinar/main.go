/*
Треба змоделювати як сигнал від мозку надходить до кінцівок людини, якщо та не танцює (достатньо ніг), але, припустимо,
у випадку, коли людина сідає, сигнал припиняє надходити.
А як знов починає танцювати - сигнал знову починає надходити. Треба опрацювати закриття каналу і відключення воркерів,
які читали з цього каналу. І знову запуск читання з каналу воркерами.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func dancing(d int) {
	if d == 1 {
		fmt.Println("I'm dancing now")
		return
	}
	fmt.Println("I'm sitting now")
}

func main() {
	wg := new(sync.WaitGroup)
	signals := [10]int{1, 1, 0, 0, 1, 1, 1, 1, 0, 0} // 1 - танцюємо, 0 сидимо
	legsChan := make(chan int)

	wg.Add(1)
	go func(signals [10]int) {
		defer wg.Done()
		for _, signal := range signals {
			legsChan <- signal
		}
	}(signals)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < len(signals); i++ {
			dancing(<-legsChan)
			time.Sleep(1 * time.Second)
		}
	}()
	wg.Wait()
}
