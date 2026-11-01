//— Напишите программу, которая предоставляет возможность принимать пришедшие по сети текстовые сообщения и сохранять их внутри себя — Мы, как клиенты, можем отправлять в эту программу на определённый эндпоинт свои сообщения по HTTP, а программа сохранит их и выведет на экран — Необходимо протестировать написанную программу

package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var readSlice []string
var mtx sync.Mutex

func ReadHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	httpRequestBodyString := string(httpRequestBody)
	fmt.Println(httpRequestBodyString)

	mtx.Lock()
	readSlice = append(readSlice, httpRequestBodyString)
	mtx.Unlock()
	fmt.Println(readSlice)

}

func main() {
	http.HandleFunc("/", ReadHandler)

	fmt.Println("starting server")
	if err := http.ListenAndServe(":228", nil); err != nil {
		fmt.Println(err)
	}

}
