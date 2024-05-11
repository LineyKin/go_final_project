package db_creator

import (
	"fmt"
	"os"

	env "go_final_project/helpers/env"
)

const env_key string = "TODO_DBFILE"

func Create() {
	dbFileName := env.GetByKey(env_key)
	dbPath := "./" + dbFileName

	// проверяем, есть ли файл БД
	_, err := os.Stat(dbPath)
	if err != nil {
		createDbFile(dbPath) // создаём файл БД
	}

	// создаём таблицу scheduler

}

func createDbFile(dbPath string) {
	_, err := os.Create(dbPath)
	if err != nil {
		fmt.Println(err)
	}
}
