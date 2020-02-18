package environment

import (
	"os"
	"strconv"
)

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func GetIntEnv(key string, defaultVal int) int {
	str := GetEnv(key, "")
	if value, err := strconv.Atoi(str); err == nil {
		return value
	}
	return defaultVal
}