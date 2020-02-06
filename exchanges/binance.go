package exchanges

import (
	"fmt"
	"github.com/RedClusive/ccspectator/database"
	"io/ioutil"
	"log"
	"net/http"
)

type Binance struct {
	Name, Url, TPrice string
}

func (cur *Binance) Parse(b *[]byte, pairs, prices *[]string) {
	cnt := 0
	CurPair := ""
	CurPrice := ""
	for _, v := range *b {
		if string(v) == string(34) {
			cnt++
			if cnt % 8 == 3 {
				if len(CurPair) != 0 {
					*pairs = append(*pairs, CurPair)
					CurPair = ""
				}
			}
			if cnt % 8 == 7 {
				if len(CurPrice) != 0 {
					*prices = append(*prices, CurPrice)
					CurPrice = ""
				}
			}
		} else {
			if cnt % 8 == 3 {
				CurPair += string(v)
			}
			if cnt % 8 == 7 {
				CurPrice += string(v)
			}
		}
	}
	if len(CurPair) != 0 {
		*pairs = append(*pairs, CurPair)
		CurPair = ""
	}
	if len(CurPrice) != 0 {
		*prices = append(*prices, CurPrice)
		CurPrice = ""
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
	res, err := http.Get((*cur).GetUrl() + (*cur).GetQueryName())
	if err != nil {
		log.Println("Exchange:", cur.GetExchangeName())
		log.Println("Can't do query:")
		log.Println(err)
		return
	}
	ToParse, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Can't read from Body:")
		log.Fatal(err)
	}
	err = res.Body.Close()
	if err != nil {
		fmt.Println("Can't close Body:")
		log.Fatal(err)
	}
	var pairs, prices []string
	(*cur).Parse(&ToParse, &pairs, &prices)
	database.SaveInDB(&pairs, &prices, cur.GetExchangeName())
}