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

	MeasureLogs := make(map[string][]sensors.Measure)

	actsensor.ActSensorPool(ActContext, MeasureChan, 2, wg)
	humsensor.HuSensorPool(HuContext, MeasureChan, 2, wg)
	presensor.PreSensorPool(PreContext, MeasureChan, 2, wg)

	go func() {
		wg.Wait()
		close(MeasureChan)
	}()

	for v := range MeasureChan {

		MeasureLogs[v.SensorType] = append(MeasureLogs[v.SensorType], v)

		fmt.Println("данные переданы в архив")

		//time.Sleep(100 \* time.Millisecond)

	}

	for k, v := range MeasureLogs {
		fmt.Println()
		switch k {
		case "pressure":
			{
				fmt.Println("собранные данные в датчиках давления")
				for i := range v {
					fmt.Println()
					fmt.Println("координаты X: ", v[i].Data.X)
					fmt.Println("координаты Y: ", v[i].Data.Y)
					fmt.Println("собранные данные по координатам: ", v[i].Data.Value, "мм рт. ст.")

				}
			}
		case "humidity":
			fmt.Println("собранные данные в датчиках влажности")
			for i := range v {
				fmt.Println()
				fmt.Println("координаты X: ", v[i].Data.X)
				fmt.Println("координаты Y: ", v[i].Data.Y)
				fmt.Println("собранные данные по координатам: ", v[i].Data.Value, "%")

			}
		case "activity":
			fmt.Println("собранные данные в датчиках активности")
			for i := range v {
				fmt.Println()
				fmt.Println("координаты X: ", v[i].Data.X)
				fmt.Println("координаты Y: ", v[i].Data.Y)
				fmt.Println("собранные данные по координатам: ", v[i].Data.Value, "пунктов")

			}
		}
		fmt.Println()
	}

}
