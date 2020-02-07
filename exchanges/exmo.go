package exchanges

import (
	"encoding/json"
	"github.com/RedClusive/ccspectator/database"
	"io/ioutil"
	"log"
	"net/http"
)

type Exmo struct {
	Name, Url, TPrice string
}

func (cur *Exmo) Parse(resp *http.Response, pairs, prices *[]string) {
	type Ticker struct {
		Avg string
	}
	jsonStream, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(cur.GetExchangeName(), err)
	}
	var dat map[string]Ticker
	if err := json.Unmarshal(jsonStream, &dat); err != nil {
		log.Println(cur.GetExchangeName(), err)
	}
	for key, value := range dat {
		*pairs = append(*pairs, key)
		*prices = append(*prices, value.Avg)
	}
}

func (cur *Exmo) GetUrl() string {
	return cur.Url
}

func (cur *Exmo) GetQueryName() string {
	return cur.TPrice
}

func (cur *Exmo) GetExchangeName() string {
	return cur.Name
}

func (cur *Exmo) DoQuery(ch chan bool) {
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