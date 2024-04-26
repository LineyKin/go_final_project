package main

import (
	"fmt"
	"net/http"
	"time"
)

func mainHandle(res http.ResponseWriter, req *http.Request) {
	s := time.Now().Format("02.01.2006 15:04:05")
	res.Write([]byte(s))
}

func main() {
	port := "7540"
	fmt.Println("Запускаем сервер")
	http.HandleFunc(`/`, mainHandle)
	fmt.Printf("http://localhost:%s/", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Завершаем работу")
}
