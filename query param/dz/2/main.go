//— Мы модифицируем программу из второго домашнего задания 19й темы — В хендлере, который призван отдавать клиенту в HTTP ответе все сохранённые сообщения, необходимо добавить возможность, при помощи query параметров, фильтровать получаемые клиентом сообщения по следующим полям: 1. Почтовый индекс отправителя 2. Срочное ли сообщение, или же нет

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

	var postIndexQ, quicklyQ string

	isPostIndex := r.URL.Query().Has("postIndex")
	isQuickly := r.URL.Query().Has("quickly")
	if isPostIndex {
		postIndexQ = r.URL.Query().Get("postIndex")
	}

	if isQuickly {
		quicklyQ = r.URL.Query().Get("quickly")
	}

	switch {
	case (isPostIndex && isQuickly):
		{
			x := make([]Msg, 0)

			postIndex, err := strconv.Atoi(postIndexQ)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println(err)
			}

			quickly, err := strconv.ParseBool(quicklyQ)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			mtx.Lock()
			for _, v := range readMap {
				if postIndex == v.PostIndex && quickly == v.Quickly {
					x = append(x, v)
				}

			}

			defer mtx.Unlock()
			if err := json.NewEncoder(w).Encode(x); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Println(err)
				return
			}
		}
	case isQuickly:
		{
			x := make([]Msg, 0)

			quickly, err := strconv.ParseBool(quicklyQ)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			mtx.Lock()
			for _, v := range readMap {
				if quickly == v.Quickly {
					x = append(x, v)
				}

			}

			defer mtx.Unlock()
			if err := json.NewEncoder(w).Encode(x); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Println(err)
				return
			}
		}
	case isPostIndex:
		{
			x := make([]Msg, 0)

			postIndex, err := strconv.Atoi(postIndexQ)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println(err)
			}

			mtx.Lock()
			for _, v := range readMap {
				if postIndex == v.PostIndex {
					x = append(x, v)
				}

			}

			defer mtx.Unlock()
			if err := json.NewEncoder(w).Encode(x); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Println(err)
				return
			}
		}
	default:
		x := make([]Msg, 0)
		mtx.Lock()
		for _, v := range readMap {
			x = append(x, v)
		}

		defer mtx.Unlock()
		if err := json.NewEncoder(w).Encode(x); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
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
