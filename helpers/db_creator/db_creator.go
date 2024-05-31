package db_creator

import (
	"database/sql"
	"fmt"
	"os"

	env "go_final_project/helpers/env"

	_ "modernc.org/sqlite"
)

const env_key string = "TODO_DBFILE"
const driverName string = "sqlite"

func getDbPath() string {
	dbFileName := env.GetDbName()

	return "./" + dbFileName
}

func Create() {
	dbPath := getDbPath()

	// проверяем, есть ли файл БД
	_, err := os.Stat(dbPath)
	if err != nil {
		createDbFile(dbPath) // создаём файл БД, если его нет
	}

	createTable() // создаём таблицу scheduler, если её нет

}

func createDbFile(dbPath string) {
	_, err := os.Create(dbPath)
	if err != nil {
		fmt.Println(err)
	}
}

func GetConnection() (*sql.DB, error) {
	return sql.Open(driverName, getDbPath())
}

func createTable() {
	db, err := sql.Open(driverName, getDbPath())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	sql := `CREATE TABLE IF NOT EXISTS scheduler (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date CHAR(8) NOT NULL DEFAULT "",
		title VARCHAR(256) NOT NULL DEFAULT "",
		comment TEXT NOT NULL DEFAULT "",
		repeat VARCHAR(128) NOT NULL DEFAULT ""
	);
	CREATE INDEX IF NOT EXISTS scheduler_date ON scheduler (date);`

	_, err = db.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
}
