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
	tablename 		  = "ratesinfotable"
	InsertStatement   = "INSERT INTO ratesinfotable (pairname, exchangename, rate, time) WHERE NOT EXISTS (SELECT * FROM ratesinfotable WHERE pairname = $1) VALUES ($1, $2, $3, $4)"
	TruncateStatement = "TRUNCATE ratesinfotable RESTART IDENTITY"
	UpdateStatement   = "UPDATE ratesinfotable SET rate = $3, time = $4 WHERE pairname = $1 AND exchangename = $2"
	SelectStatement	  = "SELECT * FROM ratesinfotable WHERE id = $1"
	CreateStatement   = "CREATE TABLE IF NOT EXISTS ratesinfotable (id SERIAL PRIMARY KEY, pairname TEXT, exchangename TEXT, rate TEXT, time TEXT)"
	DropTable 		  = "DROP TABLE IF EXISTS ratesinfotable"
)

func ConnectToDB() *sql.DB  {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Can't open data base:")
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
	_, err := db.Exec(CreateStatement)
	if err != nil {
		fmt.Println("Can't create table:")
		log.Fatal(err)
	}
	fmt.Println("Table is ready to use!")
}

func Afterparty() {
	fmt.Print("Deleting the table...")
	db := ConnectToDB()
	defer DBClose(db)
	_, err := db.Exec(DropTable)
	if err != nil {
		fmt.Println("Can't create table:")
		log.Fatal(err)
	}
	fmt.Println("Done!")
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
		log.Println("Can't insert row: ", err)
	}
}

func SelectRow(id int, pair, exchange, rate, t *string) bool {
	db := ConnectToDB()
	defer DBClose(db)
	row := db.QueryRow(SelectStatement, id)
	err := row.Scan(&id, pair, exchange, rate, t)
	if err == sql.ErrNoRows {
		return false
	}
	if err != nil {
		log.Println("Can't select row: ", err)
	}
	return true
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
