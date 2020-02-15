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
	"sync"
	"time"
)

var Fp map[string]string

func Init(m *int, Exchanges *[]exchanges.Exchange) {
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

func UpdRates(m int, quit chan bool, exchanges *[]exchanges.Exchange) {
	for {
		select {
		case <-quit:
			return
		default:
			var wg sync.WaitGroup
			for _, cur := range *exchanges {
				wg.Add(1)
				go cur.DoQuery(&wg)
			}
			wg.Wait()
			time.Sleep(time.Second * time.Duration(m))
		}
	}
}

func GetRates() string {
	fmt.Println("Getting rates...")
	type Ticker struct {
		Pair, Exchange, Rate, Updated string
	}
	var info []Ticker
	var pair, exchange, rate, t string
	for i := 1; true ; i++ {
		if !database.SelectRow(i, &pair, &exchange, &rate, &t) {
			break
		}
		pair = Fp[pair]
		if rate != "none" {
			info = append(info, Ticker{pair, exchange, rate, t});
		} else {
			log.Println("Info: no such pair:", pair)
		}
	}
	res, err := json.Marshal(info)
	if err != nil {
		log.Println("GetRates, can't Marshal info: ", err)
	} else {
		fmt.Println("Getting rates: successfully done!")
	}
	return string(res)
}

func main() {
	exsList := []exchanges.Exchange{
		&exchanges.Binance{
			Name: "Binance",
			Url: "https://api.binance.com/api/",
			Tprice: "v3/ticker/price",
		},
		&exchanges.Exmo{
			Name: "Exmo",
			Url: "https://api.exmo.com/",
			Tprice: "v1/ticker",
		},
	}
	Fp = make(map[string]string)
	var m int
	database.PrepareDB()
	Init(&m, &exsList)
	quit := make(chan bool)
	go UpdRates(m, quit, &exsList)
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		_, err := fmt.Fprint(w, GetRates())
		if err != nil {
			log.Fatal(err)
		}
	}

	http.HandleFunc("/get_rates", h1)

	port := database.GetEnv("PORT", "8000")

	log.Fatal(http.ListenAndServe(":" + port, nil))
}