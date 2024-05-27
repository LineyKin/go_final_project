package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	dbCreator "go_final_project/helpers/db_creator"
	nd "go_final_project/helpers/next_date"
	p "go_final_project/helpers/port"
	tsk "go_final_project/models/task"

	"github.com/gin-gonic/gin"
)

const webDir = "web"

// go test -run ^TestNextDate$ ./tests - OK
func getNextDate(c *gin.Context) {
	now := c.Query("now")
	if now == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "нет параметра now"})
	}

	date := c.Query("date")
	if date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "нет параметра date"})
	}

	repeat := c.Query("repeat")

	nowTime, _ := time.Parse(nd.DateFormat, now)
	nextDate, err := nd.Calc(nowTime, date, repeat)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, nextDate)
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
// go test -run ^TestDone$ ./tests - OK
func doneTask(c *gin.Context) {

	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "нет параметра id"})
		return
	}

	result, err := tsk.Done(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// обрадотчик редактирования задачи
// go test -run ^TestEditTask$ ./tests - OK
func editTask(c *gin.Context) {

	if c.Request.Method != http.MethodPut {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	var task tsk.TaskFromDB
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка десериализации JSON"})
		return
	}

	newTaskId, err := tsk.Edit(task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": newTaskId})
}

// обработчик получения задачи по её id
// go test -run ^TestTask$ ./tests - OK
func getTask(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "нет параметра id"})
		return
	}

	// получаем задачу
	task, err := tsk.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// обработчик получения списка задач
// go test -run ^TestTasks$ ./tests - OK
func getTasks(c *gin.Context) {
	list, err := tsk.GetList()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": list})
}

// обработчик добавления задачи
// go test -run ^TestAddTask$ ./tests
func addTask(c *gin.Context) {

	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	var task tsk.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка десериализации JSON"})
		return
	}

	newTaskId, err := tsk.Add(task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": newTaskId})
}

func main() {
	// проверяем БД и в случае отсутствия создаём её с таблицей
	dbCreator.Create()

	r := gin.Default()

	// API nextdate
	r.GET("/api/nextdate", getNextDate)

	// ручка добавления задачи
	r.POST("/api/task", addTask)

	// ручка получения списка задач
	r.GET("/api/tasks", getTasks)

	// ручка получения задачи по её id
	r.GET("/api/task", getTask)

	// ручка для редактирования задачи
	r.PUT("/api/task", editTask)

	// ручка отметки о выполнении задачи
	r.POST("/api/task/done", doneTask)

	// ручка удаления задачи
	//r.Delete("/api/task", deleteTask)

	// go test -run ^TestApp$ ./tests
	// go test ./tests
	//http.Handle(`/`, http.FileServer(http.Dir(webDir)))
	//http.Handle(`/js/scripts.min.js`, http.FileServer(http.Dir(webDir)))
	//http.Handle(`/css/style.css`, http.FileServer(http.Dir(webDir)))
	//http.Handle(`/favicon.ico`, http.FileServer(http.Dir(webDir)))

	r.Static("/js", "./web/js")
	r.Static("/css", "./web/css")
	r.StaticFile("/favicon.ico", "./web/favicon.ico")

	r.GET("/", fileServerHandler)

	port := p.Get() // порт из .ENV
	fmt.Println("Запускаем сервер")
	fmt.Printf("http://localhost:%s/", port)
	err := r.Run(":" + port)
	if err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
	//err := http.ListenAndServe(":"+port, nil)
	//if err != nil {
	//	fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
	//	return
	//}
}

func fileServerHandler(c *gin.Context) {
	//webDir := config.WebDir()
	filePath := filepath.Join(webDir, strings.TrimPrefix(c.Request.URL.Path, "/"))
	c.File(filePath)
}
