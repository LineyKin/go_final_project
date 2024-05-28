package task

import (
	"net/http"

	tsk "go_final_project/models/task"

	"github.com/gin-gonic/gin"
)

// обработчик удаления задачи
func DeleteTask(c *gin.Context) {
	if c.Request.Method != http.MethodDelete {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "нет параметра id"})
		return
	}

	result, err := tsk.DeleteById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// обработчик отметки о выполнении задач
// go test -run ^TestDone$ ./tests - OK
func DoneTask(c *gin.Context) {

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
func EditTask(c *gin.Context) {

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
func GetTask(c *gin.Context) {
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
func GetTasks(c *gin.Context) {
	search := c.Query("search")

	var list []tsk.TaskFromDB
	var err error
	if search == "" {
		list, err = tsk.GetList()
	} else {
		list, err = tsk.GetListBySearch(search)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": list})
}

// обработчик добавления задачи
// go test -run ^TestAddTask$ ./tests
func AddTask(c *gin.Context) {

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
