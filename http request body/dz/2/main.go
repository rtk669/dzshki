//— Мы модифицируем программу из предыдущего задания — Отныне каждое пришедшее сообщение маркируется уникальным случайно-сгенерированным ID'шников — Наша программа, помимо возможности сохранять сообщения, теперь предоставляет возможость удалять сообщения по их ID — То есть линия взаимодействия клиента с программой такова: 1. Клиент отправляет по HTTP нам несколько сообщений, текст сообщения передаётся в теле запроса. Мы каждому пришедшему сообщению генерируем случайный ID, после чего сохраняем. 2. Клиент может отправить HTTP запрос на удаления сообщения, ID удаляемого сообщения передаётся в теле запроса. 3. Каждый раз мы выводим информацию о том, какое сообщение добавлено или удалено, а так же информацию о том, какие вообще сообщения у нас сохранены, в консоль.

package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
)

var readMap map[int]string = make(map[int]string)

func ReadHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	id := rand.Intn(1000)

	readMap[id] = httpRequestBodyString

	for k, v := range readMap {
		w.Write([]byte("  ваши значения:" + fmt.Sprint(k) + " " + v))
	}

}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	id, err := strconv.Atoi(httpRequestBodyString)

	if err != nil {
		fmt.Println(err)
		return
	}
	if _, ok := readMap[id]; !ok {
		w.Write([]byte("нет такого ключа"))
		return
	} else {
		delete(readMap, id)
		w.Write([]byte("успешное удаление по ключу"))

	}

}

func main() {
	http.HandleFunc("/read", ReadHandler)
	http.HandleFunc("/delete", DeleteHandler)

	fmt.Println("starting server")
	if err := http.ListenAndServe(":228", nil); err != nil {
		fmt.Println(err)
	}

}
