package task

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	dbCreator "go_final_project/helpers/db_creator"
	nd "go_final_project/helpers/next_date"
)

const tableName string = "scheduler"

type Task struct {
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

type TaskFromDB struct {
	Id string `json:"id"`
	Task
}

func GetList() ([]TaskFromDB, error) {
	db, err := dbCreator.GetConnection()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer db.Close()

	sql := fmt.Sprintf("SELECT * FROM %s ORDER BY date", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	list := []TaskFromDB{}
	for rows.Next() {
		task := TaskFromDB{}
		err := rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		list = append(list, task)
	}

	return list, nil
}

func Add(task Task) (string, error) {
	db, err := dbCreator.GetConnection()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer db.Close()

	// заголовок задачи должен быть обязательно
	if task.Title == "" {
		return "", fmt.Errorf("не указан заголовок задачи")
	}

	// если дата не указана - пишем сегодняшнюю
	if task.Date == "" {
		task.Date = time.Now().Format(nd.DateFormat)
	}

	// проверка на адекватность записи даты. Парсить должно без ошибок
	date, err := time.Parse(nd.DateFormat, task.Date)
	if err != nil {
		return "", err
	}

	now := time.Now()
	nowStr := now.Format(nd.DateFormat)
	now, _ = time.Parse(nd.DateFormat, nowStr)

	if date.Sub(now) < 0 {
		if task.Repeat == "" {
			task.Date = now.Format(nd.DateFormat)
		} else {
			fmt.Println(task.Date)
			task.Date, err = nd.Calc(now, task.Date, task.Repeat)
			fmt.Println(task.Date)
			if err != nil {
				return "", err
			}
		}
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
		return "", err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return "", err
	}

	return strconv.Itoa(int(id)), nil
}
