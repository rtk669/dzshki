//— Напишите программу, которая предоставляет возможность принимать пришедшие по сети текстовые сообщения и сохранять их внутри себя — Мы, как клиенты, можем отправлять в эту программу на определённый эндпоинт свои сообщения по HTTP, а программа сохранит их и выведет на экран — Необходимо протестировать написанную программу

package main

import (
	"fmt"
	"io"
	"net/http"
)

var readSlice []string

func ReadHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	httpRequestBodyString := string(httpRequestBody)
	fmt.Println(httpRequestBodyString)

	readSlice = append(readSlice, httpRequestBodyString)

	fmt.Println(readSlice)

}

func main() {
	http.HandleFunc("/", ReadHandler)

	fmt.Println("starting server")
	if err := http.ListenAndServe(":228", nil); err != nil {
		fmt.Println(err)
	}

}
