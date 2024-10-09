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

type signal int

const (
	signalSit signal = iota
	signalDance
)

func dancing(s signal) {
	switch s {
	case signalSit:
		fmt.Println("I'm sitting now")
	case signalDance:
		fmt.Println("I'm dancing now")
	}
}

func main() {
	signals := [10]signal{signalSit, signalDance, signalSit, signalSit, signalDance, signalDance, signalDance, signalDance, signalSit, signalSit}
	legsChan := make(chan signal)

	go func() {
		for _, sig := range signals {
			legsChan <- sig
			time.Sleep(1 * time.Second)
		}
		close(legsChan)
	}()
	for s := range legsChan {
		dancing(s)
	}
}
