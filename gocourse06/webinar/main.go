/*
Треба зробити спрощену модель роботи внутрішніх органів людини. Достатньо серця, легень і судин між ними.
Органи працюють незалежно один від одного, але легені направляють до серця кисень. І якщо кисню не буде —
серце зупиниться.
Код повинен щонайменше використовувати горутини, канал і конструкцію select/case.
Можна time.Sleep.
*/
package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type Oxygen struct {
	Amount int
}

func vessels(lungsChan <-chan *Oxygen, heartChan chan<- *Oxygen) {
	for {
		oxygen, ok := <-lungsChan
		if !ok {
			return
		}
		heartChan <- oxygen
	}
}

func lungs(lungsChan chan<- *Oxygen) {
	for {
		oxygen := &Oxygen{Amount: rand.IntN(100)}
		fmt.Println("Lungs received oxygen: ", oxygen.Amount)
		lungsChan <- oxygen
		time.Sleep(time.Second * 1)
	}
}

// Функція, що моделює роботу серця
func heart(heartChan <-chan *Oxygen, stopChan chan<- bool) {
	for {
		select {
		case oxygen, ok := <-heartChan:
			if !ok {
				fmt.Println("Heart stopped due to lack of oxygen")
				stopChan <- true
				return
			}
			switch {
			case oxygen.Amount > 40:
				fmt.Println("Heart is pumping with oxygen amount:", oxygen.Amount)
			case oxygen.Amount < 40 && oxygen.Amount > 10:
				fmt.Println("Heart is struggling due to low oxygen")
			case oxygen.Amount < 10:
				fmt.Println("The heart does not have enough oxygen")
				stopChan <- true
				return
			}
		case <-time.After(2 * time.Second): // Якщо немає кисню більше N секунд
			fmt.Println("Heart stopped due to no oxygen supply")
			stopChan <- true
			return
		}
	}
}

func main() {
	lungsChan := make(chan *Oxygen)
	heartChan := make(chan *Oxygen)
	stopChan := make(chan bool)

	go lungs(lungsChan)
	go vessels(lungsChan, heartChan)
	go heart(heartChan, stopChan)

	<-stopChan
}
