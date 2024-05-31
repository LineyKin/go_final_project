package main

import (
	"fmt"

	hand_fs "go_final_project/handlers/fileserver"
	hand_nd "go_final_project/handlers/next_date"
	hand_tsk "go_final_project/handlers/task"
	dbCreator "go_final_project/helpers/db_creator"
	"go_final_project/helpers/env"

	"github.com/gin-gonic/gin"
)

func main() {
	// проверяем БД и в случае отсутствия создаём её с таблицей
	dbCreator.Create()

	r := gin.Default()

	// API nextdate
	r.GET("/api/nextdate", hand_nd.GetNextDate)

	// ручка добавления задачи
	r.POST("/api/task", hand_tsk.AddTask)

	// ручка получения списка задач
	r.GET("/api/tasks", hand_tsk.GetTasks)

	// ручка получения задачи по её id
	r.GET("/api/task", hand_tsk.GetTask)

	// ручка для редактирования задачи
	r.PUT("/api/task", hand_tsk.EditTask)

	// ручка отметки о выполнении задачи
	r.POST("/api/task/done", hand_tsk.DoneTask)

	// ручка удаления задачи
	r.DELETE("/api/task", hand_tsk.DeleteTask)

	// go test -count=1 -run ^TestApp$ ./tests
	// go test ./tests
	r.Static("/js", "./web/js")
	r.Static("/css", "./web/css")
	r.StaticFile("/favicon.ico", "./web/favicon.ico")
	r.StaticFile("/index.html", "./web/index.html")
	r.StaticFile("/login.html", "./web/login.html")

	r.GET("/", hand_fs.FileServerHandler)

	port := env.GetPort() // порт из .ENV
	fmt.Println("Запускаем сервер")
	fmt.Printf("http://localhost:%s/", port)
	err := r.Run(":" + port)
	if err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}
