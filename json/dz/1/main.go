//— Напишите приложение, у которого есть 2 хендлера: 1. Первый хендлер вычитывает из входящего тела запроса JSON, представляющий пользователя со следующими полями: - Имя - Адрес проживания - Возраст - Находится ли в браке - Рост, представленный дробным числом После успешного прочтения пользователя из JSON'а в структуру языка Golang, необходимо вывести на экран получившуюся переменную
//2. Второй хендлер возвращает в теле ответа JSON, представляющий заранее заготовленную структуру пользователя (описание структуры то же, что и в случае с первым хендлером)

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Person struct {
	Name    string  `json:"name"`
	City    string  `json:"city"`
	Age     int     `json:"age"`
	Married bool    `json:"married"`
	Height  float64 `json:"height"`
}

func ReqHandler(w http.ResponseWriter, r *http.Request) {
	var person Person
	// json.NewDecoder(r.Body).Decode(&person)

	// fmt.Println(person)

	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	if err := json.Unmarshal(b, &person); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	fmt.Println(person)
}

func ResHandler(w http.ResponseWriter, r *http.Request) {
	person := Person{
		Name:    "Bob",
		City:    "Bebrinsk",
		Age:     18,
		Married: false,
		Height:  1.71,
	}

	b, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(b)
}

func main() {
	http.HandleFunc("/req", ReqHandler)
	http.HandleFunc("/res", ResHandler)

	fmt.Println("starting server")
	if err := http.ListenAndServe(":228", nil); err != nil {
		fmt.Println(err)
	}

}
