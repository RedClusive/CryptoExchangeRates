package database

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

func SetUpConfig() {
	host = GetEnv("DBHOST", host)
	port= GetIntEnv("DBPORT", port)
	user = GetEnv("DBUSER", user)
	password = GetEnv("DBPASSWORD", password)
	dbname = GetEnv("DBNAME", dbname)
	db_url = GetEnv("DATABASE_URL", db_url)
}