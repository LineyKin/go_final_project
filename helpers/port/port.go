package port

import (
	env "go_final_project/helpers/env"
)

const env_key string = "TODO_PORT"

func Get() string {
	return env.GetByKey(env_key)
}
