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
	HTTPS_enabled                       bool
	Payload                             string

	Config Configure
}

func New() *HttpHeader {
	make_headers := make(map[string]string)
	return &HttpHeader{Headers: make_headers, HTTPS_enabled: true}
}

func (m *HttpHeader) PrintHeaders() *HttpHeader {
	for k, v := range m.Headers {
		fmt.Println(k, ":", v)
		fmt.Println("-------")
	}

	return m
}

func (m *HttpHeader) PrintFormData() *HttpHeader {
	fmt.Println(m.PostData)
	return m
}
