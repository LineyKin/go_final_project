package helpers

import (
	"os"

	env "github.com/joho/godotenv"
)

func GetPort() string {
	err := env.Load(".ENV")
	if err != nil {
		panic("Невозможно загрузить .ENV")
	}

	return os.Getenv("TODO_PORT")
}
