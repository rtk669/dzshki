//— Опишите любой простой хендлер, который бы обрабатывал запросы только на определённом HTTP методе (метод выберите самостоятельно, выбор случаен) — Если клиент делает запрос с неверно выбранным HTTP методом, необходимо вывести к консоль сообщение об ошибке, а статус код должен быть 405 Method Not Allowed. — Если клиент делает запрос с верно выбранным HTTP методом, то необходимо вывести в консоль информацию о том, что HTTP хендлер отработал успешно.

package main

import (
	"fmt"
	"net/http"
)

func MethodHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	} else {
		fmt.Println("going well")
		w.Write([]byte("good method"))
	}

}

func main() {
	http.HandleFunc("/test", MethodHandler)

	fmt.Println("starting server")
	http.ListenAndServe(":228", nil)
}
