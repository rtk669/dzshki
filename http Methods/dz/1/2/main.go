//— Мы модифицируем программу из второго домашнего задания 15-й темы — Помимо добавления и удаления сообщений, необоходимо реализовать хендлер, на который можно сделать запрос, и в HTTP ответе получить список всех имеющихся на данных момент сообщений — Также необходимо реализовать хендлер, который в теле запроса принимает ID сообщения, а в теле ответа отдаёт текст сообщения, которому принадлежит переданный ID. Если сообщения с переданным ID нет, в HTTP тело ответа необходимо положить информацию об ошибке, в HTTP статус код необходимо положить 404 Not Found

package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var readMap map[int]string = make(map[int]string)
var mtx sync.Mutex

func ReadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	id := rand.Intn(1000)

	mtx.Lock()
	if _, ok := readMap[id]; !ok {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("успешное создание"))
		readMap[id] = httpRequestBodyString
	} else {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("повторите попытку"))
		return
	}

	for k, v := range readMap {
		fmt.Println(("ваши значения:" + fmt.Sprint(k) + " " + v))
	}
	mtx.Unlock()

}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	id, err := strconv.Atoi(httpRequestBodyString)

	if err != nil {
		fmt.Println(err)
		return
	}
	mtx.Lock()
	if _, ok := readMap[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("нет такого ключа"))
		return
	} else {
		delete(readMap, id)
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("успешное удаление по ключу"))

	}
	mtx.Unlock()

}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	mtx.Lock()
	for k, v := range readMap {
		w.Write([]byte("  ваши значения:" + fmt.Sprint(k) + " " + v))
	}
	mtx.Unlock()
}

func main() {
	http.HandleFunc("/read", ReadHandler)
	http.HandleFunc("/delete", DeleteHandler)
	http.HandleFunc("/get", GetHandler)

	fmt.Println("starting server")
	if err := http.ListenAndServe(":228", nil); err != nil {
		fmt.Println(err)
	}

}
