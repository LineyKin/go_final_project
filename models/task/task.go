package task

import (
	"database/sql"
	"fmt"
	dbCreator "go_final_project/helpers/db_creator"
)

const tableName string = "scheduler"

type Task struct {
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func Add(task Task) (int, error) {
	db, err := dbCreator.GetConnection()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer db.Close()

	if task.Title == "" {
		return 0, fmt.Errorf("поле Задача не заполнено")
	}

	sqlPattern := `INSERT INTO %s (date, title, comment, repeat) VALUES(:date, :title, :comment, :repeat)`
	sqlPattern = fmt.Sprintf(sqlPattern, tableName)
	res, err := db.Exec(
		sqlPattern,
		sql.Named("tab", tableName),
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
