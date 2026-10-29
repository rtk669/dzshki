package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	str := "hello World"
	b := []byte(str)
	_, err := w.Write(b)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("all is good")
	}

}

func CancelHandler(w http.ResponseWriter, r *http.Request) {
	str := "payment canceling"
	b := []byte(str)
	_, err := w.Write(b)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("canceled")
	}

}

func PayHandler(w http.ResponseWriter, r *http.Request) {
	str := "trying to pay"
	b := []byte(str)
	_, err := w.Write(b)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("payment succeed")
	}

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/cancel", CancelHandler)
	http.HandleFunc("/pay", PayHandler)

	fmt.Println("starting server")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println(err)
	}
}
