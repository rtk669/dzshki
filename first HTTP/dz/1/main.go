package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hi, cheliksos"))
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("vse ok")
	}
}

func main() {
	http.HandleFunc("/", HelloHandler)

	fmt.Println("starting server")

	if err := http.ListenAndServe(":228", nil); err != nil {
		fmt.Println(err.Error())
	}
}
