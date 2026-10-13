/*
— Опишите функцию, которая генерирует одно случайно число, после чего пишет это число в переданный ей канал — Запустите описанную функцию сразу в нескольких горутинах, после чего в main функции прочитайте все значения, приходящие в канал — Выведите все полученные значения на экран — Убедитесь, что все горутины передали свои значения, и что в main горутине все эти значения были прочитаны
*/

package main

import (
	"fmt"
	"math/rand"
)

func Randomize(randonChan chan int) {
	x := rand.Intn(100)
	fmt.Println("я горутина и мое число", x)

	randonChan <- x
}

func main() {
	var x int
	randomChan := make(chan int)

	go Randomize(randomChan)
	go Randomize(randomChan)
	go Randomize(randomChan)

	x = <-randomChan
	fmt.Println(x)

	x = <-randomChan
	fmt.Println(x)

	x = <-randomChan
	fmt.Println(x)

}
