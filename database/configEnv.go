package database

import "github.com/RedClusive/ccspectator/environment"

func SetUpConfig() {
	host = environment.GetEnv("DBHOST", host)
	port= environment.GetIntEnv("DBPORT", port)
	user = environment.GetEnv("DBUSER", user)
	password = environment.GetEnv("DBPASSWORD", password)
	dbname = environment.GetEnv("DBNAME", dbname)
	db_url = environment.GetEnv("DATABASE_URL", db_url)
}