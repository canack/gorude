package gorude

import "fmt"

type Configure struct {
	Threads       int
	SleepInterval float32
}

type HttpHeader struct {
	Type, Endpoint, HttpVersion, Domain string
	Headers                             map[string]string
	PostData                            string
	HTTPSEnabled                        bool
	Payload                             string

	Config Configure
}

func New() *HttpHeader {
	makeHeaders := make(map[string]string)
	return &HttpHeader{Headers: makeHeaders, HTTPSEnabled: true}
}

func (target *HttpHeader) PrintHeaders() *HttpHeader {
	for k, v := range target.Headers {
		fmt.Println(k, ":", v)
		fmt.Println("-------")
	}

	return target
}

func (target *HttpHeader) PrintFormData() *HttpHeader {
	fmt.Println(target.PostData)
	return target
}
