package server

import (
	"fmt"
	"net/http"
)

var X int

func Get(G int) {
	X = G
}

func ServerHttp() {
	http.HandleFunc("/", handlerHttp)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

func handlerHttp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Ваша степень: %d", X*X)))
}
