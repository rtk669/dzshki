package main

import (
	"fmt"
	"net/http"
)

func QHandle(w http.ResponseWriter, r *http.Request) {
	popa := r.URL.Query().Get("popa")
	amount := r.URL.Query().Get("amount")

	fmt.Println()
	fmt.Println("popa: ", popa)
	fmt.Println("amount: ", amount)
}

func main() {
	http.HandleFunc("/q", QHandle)

	fmt.Println("starting server")
	if err := http.ListenAndServe(":228", nil); err != nil {
		fmt.Println(err)
	}
}
