/*
— Создайте 3 канала (chan int, chan string, chan float64) — Создайте 3 горутины — Каждая горутина принимает один из созданных каналов, после чего в бесконечном цикле начинает что-то в этот самый канал писать. Интервал между итерациями 300ms, 1s, 5s для каждой горутины соответственно. — В main горутине необходимо вычитывать приходящие в 3 канала значения и выводить их на экран, при этом мы не должны блокироваться на ожидании какого-то конкретного канала (а вдруг в другие каналы уже пришло значение, а мы всё ждём один конкретный канал?) — Как обычно, оценить консольный вывод, понять что происходит в программе, и сравнить это с ожиданиями
*/

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	intCh := make(chan int)
	strCh := make(chan string)
	fltCh := make(chan float64)

	go func() {
		var i int

		time.Sleep(1800 * time.Millisecond)

		for {
			i++
			intCh <- i
		}
	}()

	go func() {
		var i int

		time.Sleep(1500 * time.Millisecond)

		for {
			strCh <- "popa" + strconv.Itoa(i)
		}
	}()

	go func() {
		var i float64

		time.Sleep(3000 * time.Millisecond)

		for {
			fltCh <- i
		}
	}()

	for {
		time.Sleep(1 * time.Second)
		select {
		case number := <-intCh:
			{
				fmt.Println("я горутина инт отдаю значение: ", number)
			}
		case str := <-strCh:
			{
				fmt.Println("я горутина стринг отдаю значение: ", str)
			}
		case flt := <-fltCh:
			{
				fmt.Println("я горутина флоат отдаю значение: ", flt)
			}
		default:
			{
				fmt.Println("ничо не пришло")
			}
		}
	}

}
