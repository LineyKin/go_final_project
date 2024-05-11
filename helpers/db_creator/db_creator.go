package db_creator

import (
	"fmt"
	"os"

	env "go_final_project/helpers/env"
)

const env_key string = "TODO_DBFILE"

func Create() {
	dbFile := env.GetByKey(env_key)
	_, err := os.Stat("./" + dbFile)
	if err != nil {
		// создаём БД scheduler.db
		fmt.Println(err)
	}

	// создаём таблицу scheduler

}
