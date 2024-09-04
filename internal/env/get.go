package env

import (
	"fmt"
	"os"
)

func Get(key string) (string, error) {
	env, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("%s not available in env", key)
	}
	return env, nil
}
