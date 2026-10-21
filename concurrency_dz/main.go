// — Мы пишем симуляцию работы метеоцентра — Метеоцентр собирает показания с датчиков трёх видов:
//  1. Первый вид датчиков анализирует давление воздуха
//  2. Второй вид датчиков анализирует влажность воздуха
//  3. Третий вид датчиков анализирует сейсмическую активность
//  — Необходимо запустить симуляцию каждого вида датчиков в отдельных горутинах,
//  по одной горутине на датчик. Датчиков каждого вида может быть переменное количество.
//  — Каждый датчик, помимо прочей информации, передаёт координаты,
//  откуда передаваемая информация была получена
//  — main горутина должна уметь:
//  1. Контролировать работу датчиков
//  2. В любомой момент завершать работу одного из видов датчиков
//  3. Собирать всю поступающую от всех датчиков информацию
//  — Реализуйте предложенную симуляцию, учитывая все изученные правила корректной организации
//  работы конкуретного приложения

package main

import (
	"context"
	"fmt"
	"station/sensors"
	actsensor "station/sensors/ActSensor"
	humsensor "station/sensors/HumSensor"
	presensor "station/sensors/PreSensor"
	"sync"
	"time"
)

func main() {

	PreContext, PreCancel := context.WithTimeout(context.Background(), 5*time.Second)
	HuContext, HuCancel := context.WithTimeout(context.Background(), 5*time.Second)
	ActContext, ActCancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		PreCancel()
		HuCancel()
		ActCancel()
	}()

	MeasureChan := make(chan sensors.Measure)

	wg := &sync.WaitGroup{}

	MeasureLogs := []sensors.Measure{}

	actsensor.ActSensorPool(ActContext, MeasureChan, 200, wg)
	humsensor.HuSensorPool(HuContext, MeasureChan, 200, wg)
	presensor.PreSensorPool(PreContext, MeasureChan, 200, wg)

	go func() {
		wg.Wait()
		close(MeasureChan)
	}()

	for v := range MeasureChan {

		MeasureLogs = append(MeasureLogs, v)

		fmt.Println("данные переданы в архив")

		//time.Sleep(100 * time.Millisecond)

	}

	fmt.Println(len(MeasureLogs))
}
