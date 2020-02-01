package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
	"unicode"
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
		log.Println("Can't insert row: ", err)
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
