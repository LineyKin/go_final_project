package helpers

import (
	"log"
	"os"

	env "github.com/joho/godotenv"
)

func GetPort() string {
	err := env.Load(".ENV")
	if err != nil {
		log.Fatal("Невозможно загрузить .ENV")
	}
	return os.Getenv("TODO_PORT")
}
