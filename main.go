package main

import (
	"encoding/json"
	"fmt"
	"github.com/RedClusive/ccspectator/database"
	"github.com/RedClusive/ccspectator/exchanges"
	_ "github.com/lib/pq"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Exchange interface {
	GetExchangeName() string
	DoQuery(chan bool)
}

var Fp map[string]string

func Init(m *int, Exchanges *[]Exchange) {
	fmt.Println("Scanning input file...")
	file, err := os.Open("input.txt")
	defer func () {
		err = file.Close()
		if err != nil {
			fmt.Println("Can't close file: input.txt")
			log.Fatal(err)
		}
	}()
	if err != nil {
		fmt.Println("Can't open input.txt")
		log.Fatal(err)
	}
	_, err = fmt.Fscan(file, m)
	if err != nil {
		fmt.Println("Can't scan frequency from input.txt")
		log.Fatal(err)
	}
	s := ""
	for {
		_, err := fmt.Fscan(file, &s)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Can't scan pair from input.txt")
			log.Fatal(err)
		}
		Fp[database.FormatPair(&s)] = s
		for _, cur := range *Exchanges {
			database.InsertRow(s, cur.GetExchangeName(), "none", "none")
		}
	}
	fmt.Println("Input successfully done!")
}

func UpdRates(m int, ch, quit chan bool, Exchanges *[]Exchange) {
	defer func() {
		ch <- true
	} ()
	for {
		select {
		case <-quit:
			return
		default:
			tmpCh := make(chan bool)
			for _, cur := range *Exchanges {
				go cur.DoQuery(tmpCh)
			}
			for range *Exchanges {
				<-tmpCh
			}
			time.Sleep(time.Second * time.Duration(m))
		}
	}
}

func GetRates() string {
	fmt.Println("Getting rates...")
	result := "["
	var pair, exchange, rate, t string
	cnt := 0
	for i := 1; true ; i++ {
		if !database.SelectRow(i, &pair, &exchange, &rate, &t) {
			break
		}
		if rate != "none" {
			if cnt != 0 {
				result += ", "
			}
			pair = Fp[pair]
			out := fmt.Sprintf("{\"pair\":\"%v\", \"exchange\":\"%v\", \"rate\":\"%v\", \"updated\":\"%v\"}",
								pair, exchange, rate, t)
			result += out
			cnt++
		} else {
			log.Println("Info: no such pair:", Fp[pair])
		}
	}
	result += "]"
	if !json.Valid([]byte(result)) {
		log.Println("Error: Invalid .json generated by GetRates()")
		return "{\"errors\":[{\"code\":1337,\"message\":\"Bad .json returned.\"}]}"
	}
	fmt.Println("Getting rates: successfully done!")
	return result
}

func main() {
	Exchanges := []Exchange{
		&exchanges.Binance{
			"Binance",
			"https://api.binance.com/api/",
			"v3/ticker/price",
		},
		&exchanges.Exmo{
			"Exmo",
			"https://api.exmo.com/",
			"v1/ticker",
		},
	}
	Fp = make(map[string]string)
	var m int
	database.PrepareDB()
	Init(&m, &Exchanges)
	ch := make(chan bool)
	quit := make(chan bool)
	go UpdRates(m, ch, quit, &Exchanges)
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		_, err := fmt.Fprintf(w, "Use \"localhost/get_rates\" to get actual rates")
		if err != nil {
			log.Fatal(err)
		}
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		_, err := fmt.Fprintf(w, GetRates())
		if err != nil {
			log.Fatal(err)
		}
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/get_rates", h2)

	port := database.GetEnv("PORT", "8000")

	log.Fatal(http.ListenAndServe(":" + port, nil))
}