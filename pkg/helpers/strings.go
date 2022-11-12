package helpers

import (
	"errors"
	"os"
)

func LookupEnv(key string, fallback string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback, errors.New("key not found")
	}
	return value, nil
}
