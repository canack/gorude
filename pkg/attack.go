package gorude

import (
	"fmt"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

func (target *HttpHeader) sniper(p string) {
	var SniperHTTP HttpHeader
	SniperHTTP.Payload = p
	SniperHTTP.Headers = make(map[string]string)

	SniperHTTP.Domain = refactor(target.Domain, p)
	SniperHTTP.Endpoint = refactor(target.Endpoint, p)
	SniperHTTP.HttpVersion = refactor(target.HttpVersion, p)
	SniperHTTP.PostData = refactor(target.PostData, p)
	SniperHTTP.Type = refactor(target.Type, p)

	for key, value := range target.Headers {
		key = refactor(key, p)
		value = refactor(value, p)

		SniperHTTP.Headers[key] = value
	}

	time.Sleep(time.Second)
	SniperHTTP.HTTPS_enabled = target.HTTPS_enabled

	mainRequest(SniperHTTP)
}

var wg sync.WaitGroup

func (target *HttpHeader) Sniper(payload []string) {

	fmt.Println("Sniper attack started")

	poolsize := target.Config.Threads

	ant, _ := ants.NewPoolWithFunc(poolsize,
		func(i interface{}) {
			target.sniper(i.(string))
			wg.Done()
		})
	defer ant.Release()

	for _, payload_str := range payload {

		wg.Add(1)
		_ = ant.Invoke(string(payload_str))
		time.Sleep(time.Second * time.Duration(target.Config.SleepInterval))

	}
	wg.Wait()
}
