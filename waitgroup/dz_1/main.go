//— Опишите функцию, которая имитирует огородника, который удобряет и поливает определённую область огромного ботанического сада — Каждый огородник удобряет и поливает свою область какое-то случайное количество времени (500ms - 1000ms) — При этом количество огородников тоже случайное, ведь какая-то почва ещё могла не просохнуть до конца и не нуждаться в поливе на настоящий момент — Запустите в отдельных горутинах случайное количество огородников, дождитесь выполнения каждого (никаких time.Sleep для ожидания огородников), после чего в main горутине выведите на экран информацию об окончании выполнения программы.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Gardening(x int, wg *sync.WaitGroup) {

	defer wg.Done()

	N := rand.Intn(10)

	fmt.Println("я огородник: ", x, "мне нужно обработать: ", N+1, "участков")

	for range N + 1 {

		n := rand.Intn(6)
		n += 5
		time.Sleep(time.Duration(n*100) * time.Millisecond)

		fmt.Println("огородник ", x, "проводит работы")

	}

	fmt.Println("огородник: ", x, "завершил дела")
}
func main() {

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go Gardening(1, wg)

	wg.Add(1)
	go Gardening(2, wg)

	wg.Add(1)
	go Gardening(3, wg)

	wg.Add(1)
	go Gardening(4, wg)

	wg.Add(1)
	go Gardening(5, wg)

	wg.Wait()

	fmt.Println("main завершился!")

}
