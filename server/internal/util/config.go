package util

import (
	"fmt"
	"os"
	"strconv"
)

type Env = any

func GetEnv[T Env](key string, defaultValue T) T {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	var result Env
	switch Env(defaultValue).(type) {
	case string:
		result = value
	case int:
		v, err := strconv.Atoi(value)

		if err != nil {
			fmt.Printf("Error converting %s to int: %v\n", value, err)
			return defaultValue
		}

		result = v
	case bool:
		v, err := strconv.ParseBool(value)

		if err != nil {
			fmt.Printf("Error converting %s to int: %v\n", value, err)
			return defaultValue
		}

		result = v
	default:
		result = defaultValue
	}

	return result.(T)
}
