/*
Треба написати програму для нашийника, який одягають небезпечним тваринам: великим кішкам, великим приматам,
ведмедям та іншим. При ініціалізації нашийник відстежує пульс і температуру тварини, і з цього робить висновок,
що це за тварина. Далі записує інтенсивність дихання і звуків тварини. Якщо є сигнал від GPRS модуля — передає на
головний сервер (умовно), якщо немає, наприклад, тварина відійшла від зони покриття — записує у внутрішню памʼять,
але як сигнал зʼявиться — все накопичене передасть на головний сервер.
Де можна використати генеріки. Структура тварини передається у функцію, яка отримує дані від датчика дихання.
А потім у функцію, яка отримує дані від датчика звуку. Ця структура зберігає в собі зібрані дані та далі передає
їх на збереження (описано вище).
При реалізації треба використати канали та генеріки на рівні інтерфейсу, структури та її методів. Покрити
функції тестами.
*/
package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/sarff/prjctr-golang_beginning/gocourse08/collartask/collar"

	"github.com/sarff/prjctr-golang_beginning/gocourse08/collartask/animal"
)

func main() {
	newAnimal := animal.NewAnimal()
	fmt.Println("The collar is initialized to: ", newAnimal.Typify())
	collarDevice := collar.NewCollar()

	for range 20 {
		collarDevice.SetStrategy(&collar.GprsOn{})
		if rand.IntN(2) == 1 {
			collarDevice.SetStrategy(&collar.GprsOff{})
		}
		newAnimal.RecordBreathing(rand.Float64() * 100)
		newAnimal.RecordSound(rand.Float64() * 100)
		go collarDevice.ProcessData()
		collarDevice.Save(newAnimal)
	}
}
