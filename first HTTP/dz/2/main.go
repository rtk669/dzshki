//— Напишите 3 HTTP хендлера: 1. Первый хендлер будет возвращать "Я собака и я говорю 'Гав'" 2. Второй хендлер будет возвращать "Я кошка и я говрою 'Мяу'" 3. Третиь хендлер будет возвращать "Я корова и я говорю 'Мууу'" — Каждый хендлер должен обрабатываться на своём уникальном паттерне — Запустите HTTP сервер и протестируйте написанные хендлеры

package main

import (
	"fmt"
	"net/http"
)

func CatHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("котик говорит мяяяяяу"))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("cat cat cat")
	}
}

func DogHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("собачка говорит гав гав"))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("dog dog dog")
	}
}

func CowHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("корова говорит муууууу"))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("cow cow cow")
	}
}

func main() {
	http.HandleFunc("/cat", CatHandler)
	http.HandleFunc("/dog", DogHandler)
	http.HandleFunc("/cow", CowHandler)

	fmt.Println("starting server")
	if err := http.ListenAndServe(":228", nil); err != nil {
		fmt.Println(err.Error())
	}
}
