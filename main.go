package main

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/RedClusive/ccspectator/database"
	"github.com/RedClusive/ccspectator/environment"
	"github.com/RedClusive/ccspectator/exchanges"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"sync"
	"time"
)

var FormattedPair map[string]string = make(map[string]string)

func Init(sleepDur *time.Duration, exsList *[]exchanges.Exchange) error {
	fmt.Println("Scanning input file...")
	type Config struct {
		SleepSec int
		Pairs []string
	}
	var conf Config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		return err
	}
	*sleepDur = time.Duration(conf.SleepSec) * time.Second
	for _, s := range conf.Pairs {
		FormattedPair[database.FormatPair(&s)] = s
		for _, cur := range *exsList {
			database.InsertRow(s, cur.GetExchangeName(), "none", "none")
		}
	}
	fmt.Println("Input successfully done!")
	return nil
}

func UpdLoop(sleepDur time.Duration, quit chan struct{}, exsList *[]exchanges.Exchange) {
	for {
		select {
		case <-quit:
			return
		default:
			var wg sync.WaitGroup
			for _, cur := range *exsList {
				wg.Add(1)
				go cur.DoQuery(&wg)
			}
			wg.Wait()
			time.Sleep(sleepDur)
		}
	}
}

func GetRates() string {
	fmt.Println("Getting rates...")
	type Ticker struct {
		Pair string `json:"pair"`
		Exchange string `json:"exchange"`
		Rate string `json:"rate"`
		Updated string `json:"updated"`
	}
	var info []Ticker
	var pair, exchange, rate, t string
	for i := 1; true ; i++ {
		if !database.SelectRow(i, &pair, &exchange, &rate, &t) {
			break
		}
		pair = FormattedPair[pair]
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
			RateQuery: "v3/ticker/price",
		},
		&exchanges.Exmo{
			Name: "Exmo",
			Url: "https://api.exmo.com/",
			RateQuery: "v1/ticker",
		},
	}
	var sleepDur time.Duration
	err := database.PrepareDB()
	if err != nil {
		log.Fatal(err)
	}
	err = Init(&sleepDur, &exsList)
	if err != nil {
		log.Fatal(err)
	}
	quit := make(chan struct{})
	go UpdLoop(sleepDur, quit, &exsList)
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		_, err := fmt.Fprint(w, GetRates())
		if err != nil {
			log.Println(500, err)
		}
	}

	http.HandleFunc("/get_rates", h1)

	port := environment.GetEnv("PORT", "8000")

	log.Fatal(http.ListenAndServe(":" + port, nil))
}