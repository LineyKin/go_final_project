package helpers

import (
	"os"

	env "github.com/joho/godotenv"
)

func GetPort() string {
	env.Load(".ENV")
	return os.Getenv("TODO_PORT")
}
