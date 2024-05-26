package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	dbCreator "go_final_project/helpers/db_creator"
	nd "go_final_project/helpers/next_date"
	p "go_final_project/helpers/port"
	tsk "go_final_project/models/task"

	"github.com/go-chi/chi/v5"
)

const webDir = "web"

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
		nextDate, err = nd.Calc(nowTime, date[0], repeat[0])
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	nextDateInt, _ := strconv.Atoi(nextDate)
	resp, err := json.Marshal(nextDateInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// обработчик удаления задачи
func deleteTask(w http.ResponseWriter, r *http.Request) {
	urlStr := r.URL.String()
	urlParsed, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}

	params, _ := url.ParseQuery(urlParsed.RawQuery)

	id, idOk := params["id"]
	if !idOk {
		errorMap := map[string]string{
			"error": "отсутствует параметр id",
		}

		resp, _ := json.Marshal(errorMap)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(resp)
		return
	}

	result, err := tsk.DeleteById(id[0])
	if err != nil {
		errorMap := map[string]string{
			"error": error.Error(err),
		}

		resp, _ := json.Marshal(errorMap)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(resp)
		return
	}

	resp, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// обработчик отметки о выполнении задач
func doneTask(w http.ResponseWriter, r *http.Request) {
	urlStr := r.URL.String()
	urlParsed, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}

	params, _ := url.ParseQuery(urlParsed.RawQuery)

	id := params["id"][0]

	result, err := tsk.Done(id)
	if err != nil {
		errorMap := map[string]string{
			"error": error.Error(err),
		}

		resp, _ := json.Marshal(errorMap)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
		return
	}

	resp, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// обрадотчик редактирования задачи
func editTask(w http.ResponseWriter, r *http.Request) {
	var task tsk.TaskFromDB
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTaskId, newTaskErr := tsk.Edit(task)

	respMap := map[string]string{"id": newTaskId}
	if newTaskErr != nil {
		respMap = map[string]string{"error": error.Error(newTaskErr)}
	}

	resp, err := json.Marshal(respMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// обработчик получения задачи по её id
func getTask(w http.ResponseWriter, r *http.Request) {

	urlStr := r.URL.String()
	urlParsed, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}

	params, _ := url.ParseQuery(urlParsed.RawQuery)

	id, ok := params["id"]
	if !ok {
		respMap := map[string]string{
			"error": "нет параметра id",
		}
		resp, _ := json.Marshal(respMap)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(resp)
		return
	}

	// получаем задачу
	task, getErr := tsk.GetById(id[0])
	resp, err := json.Marshal(task)
	if getErr != nil {
		respMap := map[string]string{
			"error": error.Error(getErr),
		}
		resp, err = json.Marshal(respMap)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// обработчик получения списка задач
func getTasks(w http.ResponseWriter, r *http.Request) {
	list, listError := tsk.GetList()

	respMap := map[string][]tsk.TaskFromDB{
		"tasks": list,
	}

	resp, err := json.Marshal(respMap)

	if listError != nil {
		errorMap := map[string]string{
			"error": error.Error(listError),
		}

		resp, err = json.Marshal(errorMap)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// обработчик добавления задачи
func addTask(w http.ResponseWriter, r *http.Request) {
	var task tsk.Task
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTaskId, newTaskErr := tsk.Add(task)

	respMap := map[string]string{"id": newTaskId}
	if newTaskErr != nil {
		respMap = map[string]string{"error": error.Error(newTaskErr)}
	}

	resp, err := json.Marshal(respMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func main() {
	// проверяем БД и в случае отсутствия создаём её с таблицей
	dbCreator.Create()

	r := chi.NewRouter()

	// API nextdate
	r.Get("/api/nextdate", getNextDate)

	// ручка добавления задачи
	r.Post("/api/task", addTask)

	// ручка получения списка задач
	r.Get("/api/tasks", getTasks)

	// ручка получения задачи по её id
	r.Get("/api/task", getTask)

	// ручка для редактирования задачи
	r.Put("/api/task", editTask)

	// ручка отметки о выполнении задачи
	r.Post("/api/task/done", doneTask)

	// ручка удаления задачи
	r.Delete("/api/task", deleteTask)

	// Ручки основной страницы фронта и файлов фронта.
	// Баг: неадекватно реагирует на ctrl + shift + R,
	// но если не нажимать так, то всё ок
	fileServer(r, `/`, http.Dir(webDir))
	fileServer(r, `/js/scripts.min.js`, http.Dir(webDir))
	fileServer(r, `/css/style.css`, http.Dir(webDir))
	fileServer(r, `/favicon.ico`, http.Dir(webDir))

	//http.Handle(`/`, http.FileServer(http.Dir(webDir)))
	//http.Handle(`/js/scripts.min.js`, http.FileServer(http.Dir(webDir)))
	//http.Handle(`/css/style.css`, http.FileServer(http.Dir(webDir)))
	//http.Handle(`/favicon.ico`, http.FileServer(http.Dir(webDir)))

	port := p.Get() // порт из .ENV
	fmt.Println("Запускаем сервер")
	fmt.Printf("http://localhost:%s/", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}

// файл сервер на chi
// найдено в интернете:
// https://github.com/go-chi/chi/blob/master/_examples/fileserver/main.go
func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(path)
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
