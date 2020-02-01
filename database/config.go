package database

const (
	host 	 		  = "localhost"
	port 	 		  = 5432
	user 	 		  = "postgres"
	password 		  = "1862"
	dbname    		  = "humble_base"
	tablename 		  = "infotable"
	InsertStatement   = "INSERT INTO infotable (pairname, exchangename, rate, time) VALUES ($1, $2, $3, $4)"
	TruncateStatement = "TRUNCATE infotable RESTART IDENTITY"
	UpdateStatement   = "UPDATE infotable SET rate = $3, time = $4 WHERE pairname = $1 AND exchangename = $2"
	SelectStatement	  = "SELECT * FROM infotable WHERE id = $1"
)
