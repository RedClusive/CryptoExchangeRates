package database

import (
	"os"
	"strconv"
)

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getIntEnv(key string, defaultVal int) int {
	str := getEnv(key, "")
	if value, err := strconv.Atoi(str); err == nil {
		return value
	}
	return defaultVal
}

func SetUpConfig() {
	host = getEnv("HOST", host)
	port= getIntEnv("PORT", port)
	user = getEnv("USER", user)
	password = getEnv("DBNAME", password)
}