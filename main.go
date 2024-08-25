package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	w.Write([]byte(now.Format("02-01-2006 15:04:05")))
}

func main() {

	http.HandleFunc("/time", handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
