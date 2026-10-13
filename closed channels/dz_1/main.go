package main

func main() {
	/*

		// закрываем закрытый канал

		ch := make(chan int)

		close(ch)

		close(ch)

		//будет паника
	*/

	/*
		//читаем из закрытого канала

		ch := make(chan int)

		close(ch)

		v, ok := <-ch

		fmt.Println(v, ok)
		//читается значение по умолчанию, для проверки используем ok
	*/

	/*
		//записываем в закрытый канал

		ch := make(chan int)

		close(ch)

		ch <- 10

		//будет паника
	*/

	/*
		//закрытие nil канала

		var ch chan int

		close(ch)

		//будет паника
	*/

	/*
		//чтение из nil канала

		var ch chan int

		x := <-ch
		fmt.Println(x)

		//дедлок
	*/

	/*
		//запись в nil канал

		var ch chan int

		ch <- 10

		//будет дедлок
	*/

}
