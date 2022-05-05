package gorude

import (
	"fmt"
	"runtime"
	"sync"
	"time"
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
	SniperHTTP.HTTPSEnabled = target.HTTPSEnabled

	mainRequest(SniperHTTP)
}

var wg sync.WaitGroup

func (target *HttpHeader) Sniper(payload []string) {

	fmt.Println("Sniper attack started")

	poolSize := target.Config.Threads
	runtime.GOMAXPROCS(poolSize)

	for _, payloadString := range payload {

		wg.Add(1)

		target.sniper(payloadString)
		time.Sleep(time.Second * time.Duration(target.Config.SleepInterval))

	}
	wg.Wait()
}
