package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
	"unicode"
)

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

func ConnectToDB() *sql.DB  {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Can't open data base:")
		DBClose(db)
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func DBClose(db *sql.DB) {
	err := db.Close()
	if err != nil {
		fmt.Println("Can't close data base:")
		log.Fatal(err)
	}
}

func PrepareDB() {
	fmt.Println("Preparing the table...")

	db := ConnectToDB()
	defer DBClose(db)
	_, err := db.Exec(TruncateStatement)
	if err != nil {
		fmt.Println("Can't truncate table:")
		log.Fatal(err)
	}
	fmt.Println("Table is ready to use!")
}

func FormatPair(s *string) string {
	res := ""
	for _, c := range *s {
		if unicode.IsLetter(c) {
			res += string(c)
		}
	}
	return strings.ToUpper(res)
}

func InsertRow(pairname, exchangename, rate, time string) {
	db := ConnectToDB()
	defer DBClose(db)
	_, err := db.Exec(InsertStatement, FormatPair(&pairname), exchangename, rate, time)
	if err != nil {
		fmt.Println("Can't insert in: ", tablename)
		log.Fatal(err)
	}
}

func SaveInDB(pairs, prices *[]string, name string) {
	t := time.Now().Format("2006-01-02 15:04:05.000")
	db := ConnectToDB()
	defer DBClose(db)
	for i := range *pairs {
		_, err := db.Exec(UpdateStatement, FormatPair(&(*pairs)[i]), name, (*prices)[i], t)
		if err != nil {
			log.Println("Can't update the table:", err)
		}
	}
}
