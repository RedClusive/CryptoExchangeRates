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
	host = getEnv("DBHOST", host)
	port= getIntEnv("DBPORT", port)
	user = getEnv("DBUSER", user)
	password = getEnv("DBPASSWORD", password)
	dbname = getEnv("DBNAME", dbname)
	db_url = getEnv("DATABASE_URL", db_url)
}