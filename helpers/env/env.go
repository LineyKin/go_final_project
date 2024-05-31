package env

import (
	"os"

	env "github.com/joho/godotenv"
)

const port_key string = "TODO_PORT"
const db_key string = "TODO_DBFILE"

func getByKey(key string) string {
	err := env.Load(".ENV")

	if err != nil {
		panic("Невозможно загрузить .ENV")
	}

	return os.Getenv(key)
}

func GetPort() string {
	return getByKey(port_key)
}

func GetDbName() string {
	return getByKey(db_key)
}
