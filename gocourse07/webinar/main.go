/*
Треба змоделювати як сигнал від мозку надходить до кінцівок людини, якщо та не танцює (достатньо ніг), але, припустимо,
у випадку, коли людина сідає, сигнал припиняє надходити.
А як знов починає танцювати - сигнал знову починає надходити. Треба опрацювати закриття каналу і відключення воркерів,
які читали з цього каналу. І знову запуск читання з каналу воркерами.
*/
package main

import (
	"fmt"
	"time"
)

const (
	signalSit = iota
	signalDance
)

func dancing(d int) {
	switch d {
	case signalSit:
		fmt.Println("I'm sitting now")
	case signalDance:
		fmt.Println("I'm dancing now")
	}
}

func main() {
	// wg := new(sync.WaitGroup)
	signals := [10]int{1, 1, 0, 0, 1, 1, 1, 1, 0, 0} // 1 - танцюємо, 0 сидимо
	legsChan := make(chan int)

	go func(signals [10]int) {
		for _, signal := range signals {
			legsChan <- signal
		}
	}(signals)

	for i := 0; i < len(signals); i++ {
		dancing(<-legsChan)
		time.Sleep(1 * time.Second)
	}
}
