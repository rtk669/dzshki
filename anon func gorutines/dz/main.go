/*
— Запустите 3 горутины через различные анонимные функции — Каждая анонимная функция должна 3 раза выводить на экран любой текст — Посмотрите на консольный вывод, сравните этот вывод с вашими ожиданиями
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for range 3 {
			fmt.Println("booo")
		}

	}()

	go func() {

		for range 3 {
			fmt.Println("fooo")
		}

	}()

	go func() {
		for range 3 {
			fmt.Println("goooo")
		}

	}()

	time.Sleep(1 * time.Second)
}
