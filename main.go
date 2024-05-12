package main

import (
	"fmt"
	"net/http"
	"time"

	//dbCreator "go_final_project/helpers/db_creator"
	nextDate "go_final_project/helpers/next_date"
	p "go_final_project/helpers/port"
)

func main() {

	now := time.Now()
	date := "20240513"
	repeatRule := "d 400"

	nextDate.Get(now, date, repeatRule)

	// проверяем БД и в случае отсутствия создаём её с таблицей
	//dbCreator.Create()

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
