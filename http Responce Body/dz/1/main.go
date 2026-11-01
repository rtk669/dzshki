//— Опишите HTTP хендлер, который в зависимости от цифры, переданной в HTTP запросе, будет отдавать различные HTTP ответы и различные статус-коды

package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	req, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	reqNum, err := strconv.Atoi(string(req))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	switch reqNum {
	case 1:
		w.WriteHeader(http.StatusBadGateway)
	case 2:
		w.WriteHeader(http.StatusBadRequest)
	case 3:
		w.WriteHeader(http.StatusConflict)
	case 4:
		w.WriteHeader(http.StatusContinue)
	case 5:
		w.WriteHeader(http.StatusCreated)
	}

}

func main() {
	http.HandleFunc("/", TestHandler)

	fmt.Println("starting server")
	if err := http.ListenAndServe(":228", nil); err != nil {
		fmt.Println(err)
	}

}
