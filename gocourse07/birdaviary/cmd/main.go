/*
Вам потрібно створити програму для управління вольєром для екзотичних птахів у розумному зоопарку. Вольєр має сенсори,
які вимірюють температуру, яскравість освітлення та вологість, і вони передають ці дані до центральної системи, яка,
в свою чергу, зберігає їх у памʼять (така от база даних). Раз на добу сенсори відключаються для технічної перевірки,
але потім знову продовжують працювати. Центральна система також перезавантажується раз на добу.
Робота кожного сенсора — окрема горутина. Коли сенсори відключаються, горутини безпечно вимикається. Те саме з
центральною системою. Процес запису в памʼять повинен бути (штучно) тривалим. І у випадку, коли центральна система
планово відключається, вона це мусить робити лише після того, як всі записи в базу виконались.
*/
package main

import (
	"os"
	"sync"

	"github.com/sarff/prjctr-golang_beginning/gocourse07/birdaviary/internal/centralsystem"
	"github.com/sarff/prjctr-golang_beginning/gocourse07/birdaviary/internal/logger"
	"github.com/sarff/prjctr-golang_beginning/gocourse07/birdaviary/internal/sensor"
	"github.com/sarff/prjctr-golang_beginning/gocourse07/birdaviary/internal/storage"
)

func main() {
	wg := new(sync.WaitGroup)
	log := logger.New(os.Stdout)
	db := storage.NewStorage(log)
	cs := centralsystem.NewCentralSystem(db, log)

	var tempSensor, humiditySensor, brightnessSensor sensor.Sensor

	tempSensor = sensor.NewSensor(cs, log, "TemperatureSensor")
	humiditySensor = sensor.NewSensor(cs, log, "HumiditySensor")
	brightnessSensor = sensor.NewSensor(cs, log, "BrightnessSensor")

	wg.Add(1)
	go tempSensor.Start(wg)
	wg.Add(1)
	go humiditySensor.Start(wg)
	wg.Add(1)
	go brightnessSensor.Start(wg)

	wg.Add(1)
	go cs.Start(wg)
	wg.Wait()

	loadName := "TemperatureSensor"
	load, err := db.Load(loadName)
	if err != nil {
		log.Warn("Cant load by Name from database", "name", loadName, "error description", err)
		return
	}
	log.Info("History data:", loadName, load)
}
