package exchanges

import "sync"

type Exchange interface {
	GetExchangeName() string
	DoQuery(group *sync.WaitGroup)
}
