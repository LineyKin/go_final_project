package env

import (
	"os"

	env "github.com/joho/godotenv"
)

func GetByKey(key string) string {
	err := env.Load(".ENV")

	if err != nil {
		panic("Невозможно загрузить .ENV")
	}

	return os.Getenv(key)
}
