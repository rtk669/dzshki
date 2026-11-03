//— Мы модифицируем программу из второго домашнего задания 18й темы — Теперь каждое передаваемое HTTP клиентом в теле запроса сообщение (во время отправки запроса на создание этого самого сообщения) должно быть оформлено в теле запроса в виде JSON'а и содержать следующие поля: 1. Заголовок сообщения 2. Почтовый индекс отправителя, представленный целым числом 3. Текст сообщения 4. Срочное ли это сообщение, или же нет — Когда мы возвращаем клиенту список всех существующих сообщений, этот самый список также должен быть оформлен в виде JSON-массива в теле ответа

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

type Msg struct {
	Title     string `json:"title"`
	PostIndex int    `json:"postIndex"`
	Message   string `json:"message"`
	Quickly   bool   `json:"quickly"`
}

var readMap map[int]Msg = make(map[int]Msg)
var mtx sync.Mutex

func ReadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var msg Msg
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	id := rand.Intn(1000)

	mtx.Lock()
	defer mtx.Unlock()
	if _, ok := readMap[id]; !ok {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("успешное создание"))
		readMap[id] = msg
	} else {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("повторите попытку"))
		return
	}

	for k, v := range readMap {
		fmt.Println(" ваши значения: ", k, "  ", v)
	}

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
	defer mtx.Unlock()
	if _, ok := readMap[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("нет такого ключа"))
		return
	} else {
		delete(readMap, id)
		w.WriteHeader(http.StatusNoContent)
	}

}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	x := make([]Msg, 0)

	for _, v := range readMap {
		x = append(x, v)
	}
	mtx.Lock()
	defer mtx.Unlock()
	if err := json.NewEncoder(w).Encode(x); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

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
