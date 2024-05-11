package main

import (
	"fmt"
	"net/http"

	db_creator "go_final_project/helpers/db_creator"
	p "go_final_project/helpers/port"
)

func main() {
	db_creator.Create()
	port := p.Get()
	fmt.Println("Запускаем сервер")

	// "ручки" основной страницы фронта и файлов фронта
	webDir := "web"
	http.Handle(`/`, http.FileServer(http.Dir(webDir)))
	http.Handle(`/js/scripts.min.js`, http.FileServer(http.Dir(webDir)))
	http.Handle(`/css/style.css`, http.FileServer(http.Dir(webDir)))
	http.Handle(`/favicon.ico`, http.FileServer(http.Dir(webDir)))

	fmt.Printf("http://localhost:%s/", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Завершаем работу")
}
