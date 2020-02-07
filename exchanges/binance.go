package exchanges

import (
	"encoding/json"
	"github.com/RedClusive/ccspectator/database"
	"io"
	"log"
	"net/http"
)

type Binance struct {
	Name, Url, TPrice string
}

func (cur *Binance) Parse(resp *http.Response, pairs, prices *[]string) {
	type Ticker struct {
		Symbol, Price string
	}
	dec := json.NewDecoder(resp.Body)
	var a Ticker
	for {
		if err := dec.Decode(&a); err == io.EOF {
			break;
		} else if err != nil {
			log.Println(cur.GetExchangeName(), err)
		}
		*pairs = append(*pairs, a.Symbol)
		*prices = append(*prices, a.Price)
	}
}

func (cur *Binance) GetUrl() string {
	return cur.Url
}

func (cur *Binance) GetQueryName() string {
	return cur.TPrice
}

func (cur *Binance) GetExchangeName() string {
	return cur.Name
}

func (cur *Binance) DoQuery(ch chan bool) {
	defer func(){
		ch <- true
	}()
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
	(*cur).Parse(resp, &pairs, &prices)
	database.SaveInDB(&pairs, &prices, cur.GetExchangeName())
}