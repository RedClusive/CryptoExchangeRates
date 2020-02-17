package exchanges

import (
	"encoding/json"
	"github.com/RedClusive/ccspectator/database"
	"log"
	"net/http"
	"sync"
)

type Binance struct {
	Name, Url, Tprice string
}

func (cur *Binance) Parse(resp *http.Response, pairs, prices *[]string) {
	type Ticker struct {
		Symbol, Price string
	}
	var dec []Ticker
	err := json.NewDecoder(resp.Body).Decode(&dec)
	if err != nil {
		log.Println(cur.GetExchangeName(), err)
	}
	for _, v := range dec {
		*pairs = append(*pairs, v.Symbol)
		*prices = append(*prices, v.Price)
	}
}

func (cur *Binance) GetUrl() string {
	return cur.Url
}

func (cur *Binance) GetQueryName() string {
	return cur.Tprice
}

func (cur *Binance) GetExchangeName() string {
	return cur.Name
}

func (cur *Binance) DoQuery(wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get((*cur).GetUrl() + (*cur).GetQueryName())
	defer func(){
		err := resp.Body.Close()
		if err != nil {
			log.Println(cur.GetExchangeName(), err)
		}
	} ()
	if err != nil {
		log.Println(cur.GetExchangeName(), err)
		return
	}
	var pairs, prices []string
	cur.Parse(resp, &pairs, &prices)
	database.SaveInDB(&pairs, &prices, cur.GetExchangeName())
}