package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	dbCreator "go_final_project/helpers/db_creator"

	nd "go_final_project/helpers/next_date"
	p "go_final_project/helpers/port"
)

func getNextDate(w http.ResponseWriter, r *http.Request) {
	urlStr := r.URL.String()
	urlParsed, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}

	params, _ := url.ParseQuery(urlParsed.RawQuery)

	now, okNow := params["now"]
	if !okNow {
		http.Error(w, "нет параметра now", http.StatusBadRequest)
	}

	date, okDate := params["date"]
	if !okDate {
		http.Error(w, "нет параметра date", http.StatusBadRequest)
	}

	repeat, okRepeat := params["repeat"]
	if !okRepeat {
		http.Error(w, "нет параметра repeat", http.StatusBadRequest)
	}

	nextDate := ""
	if okNow && okDate && okRepeat {
		nowTime, _ := time.Parse(nd.DateFormat, now[0])
		nextDate, _ = nd.Get(nowTime, date[0], repeat[0])
	}

	resp, err := json.Marshal(nextDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func main() {
	// проверяем БД и в случае отсутствия создаём её с таблицей
	dbCreator.Create()

	// "ручки" основной страницы фронта и файлов фронта
	webDir := "web"
	http.Handle(`/`, http.FileServer(http.Dir(webDir)))
	http.Handle(`/js/scripts.min.js`, http.FileServer(http.Dir(webDir)))
	http.Handle(`/css/style.css`, http.FileServer(http.Dir(webDir)))
	http.Handle(`/favicon.ico`, http.FileServer(http.Dir(webDir)))

	http.Get("/api/nextdate/")
	http.HandleFunc("/api/nextdate/", getNextDate)

	port := p.Get() // порт из .ENV
	fmt.Println("Запускаем сервер")
	fmt.Printf("http://localhost:%s/", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
